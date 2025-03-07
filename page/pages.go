package page

import (
	"embed"
	"github.com/busy-cloud/boat/web"
	"io/fs"
	"net/http"
)

var pages web.Store

func Web(fs http.FileSystem, base string) {
	pages.Web(fs, base)
}

func FS(fs fs.FS, base string) {
	pages.FS(fs, base)
}

func Dir(dir string, base string) {
	pages.Dir(dir, base)
}

func Zip(zip string, base string) {
	pages.Zip(zip, base)
}

func EmbedFS(fs embed.FS, base string) {
	pages.EmbedFS(fs, base)
}

func Open(name string) (file http.File, err error) {
	return pages.Open(name)
}
