package server

import (
	"embed"
	"io/fs"
)

//go:embed frontend/public
var frontendFS embed.FS
var subFrontendFS fs.FS

func init() {
	var err error
	subFrontendFS, err = fs.Sub(frontendFS, "frontend/public")
	panicIfErr(err)
}

func ReadResource(name string) ([]byte, error) {
	return fs.ReadFile(subFrontendFS, name)
}
