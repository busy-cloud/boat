package apps

import (
	"github.com/busy-cloud/boat/store"
)

var assets store.Store

func Assets() *store.Store {
	return &assets
}

var pages store.Store

func Pages() *store.Store {
	return &pages
}

//func init() {
//	assets.AddDir("assets") //当前目录
//	pages.AddDir("pages")
//}
