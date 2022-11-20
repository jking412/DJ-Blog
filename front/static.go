package front

import (
	"embed"
	_ "embed"
	"io/fs"
)

type UI struct {
	FS         embed.FS
	StaticPath string
}

func (ui *UI) Open(path string) (fs.File, error) {
	return ui.FS.Open(ui.StaticPath + path)
}

//go:embed build
var Build embed.FS
