package domain

import "errors"

var (
	ErrDocumentNotFound  = errors.New("document not found")
	ErrVersionConflict   = errors.New("document version conflict")
	ErrWorkspaceNotFound = errors.New("workspace not found")
)
