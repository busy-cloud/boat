package page

import (
	"embed"
	"github.com/busy-cloud/boat/store"
	"net/http"
)

var pages store.Store

func Dir(dir string, base string) {
	pages.Dir(dir, base)
}

func Zip(zip string, base string) {
	pages.Zip(zip, base)
}

func EmbedFS(fs embed.FS, base string) {
	pages.EmbedFS(fs, base)
}

func Open(name string) (http.File, error) {
	file, err := pages.Open(name)
	return store.HttpFile(file), err
}
