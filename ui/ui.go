package ui

import "embed"

//go:embed "templates" "assets" "static"
var EmbeddedFiles embed.FS
