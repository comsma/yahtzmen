package web

import "embed"

var (
	//go:embed frontend/dist
	Web embed.FS
)
