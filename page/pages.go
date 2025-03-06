package page

import "github.com/busy-cloud/boat/lib"

var pages lib.Map[Page]

func Register(name string, menu *Page) {
	pages.Store(name, menu)
}

func Unregister(name string) {
	pages.Delete(name)
}
