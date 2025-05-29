package apps

import "github.com/busy-cloud/boat/store"

var pages store.Store

func Pages() *store.Store {
	return &pages
}
