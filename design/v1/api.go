package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("kollap", func() {
	Title("Collaborative Knowledge Platform Backenk service")
	Description(`
	API for backend service powering Kollap platform.
	`)
	Version("1.0")
})

var Workspace = Type("Workspace", func() {
	Description("Workspace is where multiple people can come and share knowledge using documents.")

	Attribute("id", func() {
		Format(FormatUUID)
	})
	Attribute("name", String)
	Attribute("owner_id", func() {
		Format(FormatUUID)
	})
	Attribute("created_at", func() {
		Format(FormatDateTime)
	})

	Required("id", "name", "owner_id", "created_at")
})

var Document = Type("Document", func() {
	Description("Knowledge sharing medium in this platform is a document.")

	Attribute("id", func() {
		Format(FormatUUID)
	})
	Attribute("title", String)
	Attribute("version", Int64)
	Attribute("content", String)
	Attribute("workspace_id", func() {
		Format(FormatUUID)
	})
	Attribute("created_at", String, func() {
		Format(FormatDateTime)
	})

	Required("id", "title", "version", "content", "workspace_id", "created_at")
})

var _ = Service("workspaces", func() {
	Description("Workspaces related requests")
	HTTP(func() {
		Path("/v1")
	})

	Method("createWorkspace", func() {
		Payload(func() {
			Attribute("name", String)
			Attribute("owner_id", func() {
				Format(FormatUUID)
			})

			Required("name", "owner_id")
		})
		Result(Workspace)

		HTTP(func() {
			POST("/workspaces")
		})
	})

	Method("getWorkspace", func() {
		Payload(func() {
			Attribute("id", func() {
				Format(FormatUUID)
			})

			Required("id")
		})
		Result(Workspace)

		HTTP(func() {
			GET("/workspaces/{id}")
		})
	})

	Method("getAllWorkspaceDocuments", func() {
		Payload(func() {
			Attribute("id", func() {
				Format(FormatUUID)
			})

			Required("id")
		})
		Result(func() {
			Attribute("documents", ArrayOf(Document))
		})

		HTTP(func() {
			GET("/workspaces/{id}/documents")
		})
	})
})

var _ = Service("documents", func() {
	Description("Documents related Endpoints")
	HTTP(func() {
		Path("/v1")
	})

	Method("getDocument", func() {
		Payload(func() {
			Attribute("id", func() {
				Format(FormatUUID)
			})

			Required("id")
		})
		Result(Document)

		HTTP(func() {
			GET("/documents/{id}")
		})
	})

	Method("createDocument", func() {
		Payload(func() {
			Attribute("title", String)
			Attribute("workspace_id", func() {
				Format(FormatUUID)
			})
			Attribute("content", String)

			Required("title", "workspace_id")
		})
		Result(Document)

		HTTP(func() {
			POST("/documents")
		})
	})

	Method("updateDocument", func() {
		Payload(func() {
			Attribute("id", func() {
				Format(FormatUUID)
			})
			Attribute("title", String)
			Attribute("content", String)

			Required("id")
		})
		Result(Document)

		HTTP(func() {
			PATCH("/documents/{id}")
		})
	})
})
