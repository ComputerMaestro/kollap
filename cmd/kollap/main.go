package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	documents "github.com/ComputerMaestro/kollap/gen/documents"
	workspaces "github.com/ComputerMaestro/kollap/gen/workspaces"
	httpadapters "github.com/ComputerMaestro/kollap/internal/adapters/http"
	"github.com/ComputerMaestro/kollap/internal/application/document"
	"github.com/ComputerMaestro/kollap/internal/application/workspace"
	"github.com/ComputerMaestro/kollap/internal/config"
	postgresrepo "github.com/ComputerMaestro/kollap/internal/infrastructure/postgres"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load Configuration
	conf := config.GetConfig()

	// Setup logger. Replace logger with your own log package of choice.
	format := log.FormatJSON
	if log.IsTerminal() {
		format = log.FormatTerminal
	}
	ctx := log.Context(context.Background(), log.WithFormat(format))
	if conf.Server.Debug {
		ctx = log.Context(ctx, log.WithDebug())
		log.Debugf(ctx, "debug logs enabled")
	}
	log.Print(ctx, log.KV{K: "http-port", V: conf.Server.Port})

	// DB initialization
	dsn := "host=" + conf.Db.Host +
		" user=" + conf.Db.User +
		" password=" + conf.Db.Password +
		" dbname=" + conf.Db.DBName +
		" port=" + conf.Db.Port +
		" sslmode=" + conf.Db.SSLMode +
		" TimeZone=" + conf.Db.TimeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to establish connection to db")
	}

	workspaceRepo := postgresrepo.NewWorkspacePostgresRepository(db)
	documentRepo := postgresrepo.NewDocumentPostgresRepository(db)

	createWorkspaceUC := workspace.NewCreateWorkspaceUC(workspaceRepo)
	getWorkspaceUC := workspace.NewGetWorkspaceUC(workspaceRepo)
	getAllWorkspaceDocumentsUC := workspace.NewGetAllWorkspaceDocumentsUC(documentRepo)

	createDocumentUC := document.NewCreateDocumentUC(documentRepo)
	getDocumentUC := document.NewGetDocumentUC(documentRepo)
	updateDocumentUC := document.NewUpdateDocumentUC(documentRepo)

	// Initialize the services.
	var (
		workspacesSvc workspaces.Service
		documentsSvc  documents.Service
	)
	{
		workspacesSvc = httpadapters.NewWorkspaces(
			createWorkspaceUC,
			getWorkspaceUC,
			getAllWorkspaceDocumentsUC,
		)
		documentsSvc = httpadapters.NewDocuments(
			createDocumentUC,
			getDocumentUC,
			updateDocumentUC,
		)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		workspacesEndpoints *workspaces.Endpoints
		documentsEndpoints  *documents.Endpoints
	)
	{
		workspacesEndpoints = workspaces.NewEndpoints(workspacesSvc)
		workspacesEndpoints.Use(debug.LogPayloads())
		workspacesEndpoints.Use(log.Endpoint)
		documentsEndpoints = documents.NewEndpoints(documentsSvc)
		documentsEndpoints.Use(debug.LogPayloads())
		documentsEndpoints.Use(log.Endpoint)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)

	// Start the servers and send errors (if any) to the error channel.
	switch conf.Server.Host {
	case "localhost":
		{
			addr := fmt.Sprintf("http://%s:%s", conf.Server.Host, conf.Server.Port)
			u, err := url.Parse(addr)
			if err != nil {
				log.Fatalf(ctx, err, "invalid URL %#v\n", addr)
			}
			if conf.Server.Port != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					log.Fatalf(ctx, err, "invalid URL %#v\n", u.Host)
				}
				u.Host = net.JoinHostPort(h, conf.Server.Port)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, workspacesEndpoints, documentsEndpoints, &wg, errc, conf.Server.Debug)
		}

	default:
		log.Fatal(ctx, fmt.Errorf("invalid host argument: %q (valid hosts: localhost)", conf.Server.Host))
	}

	// Wait for signal.
	log.Printf(ctx, "exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	log.Printf(ctx, "exited")
}
