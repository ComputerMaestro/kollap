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
	Attribute("created_at", String, func() {
		Format(FormatDateTime)
	})
})

var _ = Service("workspaces", func() {
	Description("Workspaces related requests")
	HTTP(func() {
		Path("/v1/workspace")
	})

	Method("createWorkspace", func() {
		Payload(func() {
			Attribute("name", String)

			Required("name")
		})
		Result(Workspace)

		HTTP(func() {
			POST("/create")
		})
	})
})

var _ = Service("documents", func() {
	Description("Documents related Endpoints")
	HTTP(func() {
		Path("/v1/document")
	})

	Method("getDocument", func() {
		Payload(func() {
			Attribute("id", func() {
				Format(FormatUUID)
			})
		})
		Result(Document)

		HTTP(func() {
			GET("/{id}")
		})
	})
})
