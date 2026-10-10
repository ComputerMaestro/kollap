package genkit

import (
	"context"
	"fmt"

	"github.com/ComputerMaestro/kollap/internal/config"
	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
	"goa.design/clue/log"
)

type GenkitEmbedder struct {
	genkitInst *genkit.Genkit
	embedder   ai.Embedder
}

func NewGenkitEmbedder(ctx context.Context, conf *config.Embedder) (*GenkitEmbedder, error) {
	o := &ollama.Ollama{
		ServerAddress: conf.ServerAddress, // Default Ollama server
		Timeout:       int(conf.Timeout),  // Response timeout in seconds
	}
	fmt.Printf("%v", conf)

	g := genkit.Init(ctx, genkit.WithPlugins(o))

	embedder := o.DefineEmbedder(g, conf.Model, int(conf.Dimensions), nil)

	return &GenkitEmbedder{
		genkitInst: g,
		embedder:   embedder,
	}, nil
}

func (g *GenkitEmbedder) Embed(ctx context.Context, doc *domain.Document) ([]float32, error) {
	res, err := genkit.Embed(ctx,
		g.genkitInst,
		ai.WithEmbedder(g.embedder),
		ai.WithDocs(ai.DocumentFromText(doc.Content, map[string]any{
			domain.DOCUMENT_TITLE: doc.Title,
		})),
	)
	if err != nil || len(res.Embeddings) == 0 {
		log.Errorf(ctx, err, "failed to generate embedding for document %v", doc.ID)
		return nil, err
	}
	return res.Embeddings[0].Embedding, nil
}
