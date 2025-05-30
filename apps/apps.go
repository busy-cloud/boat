package apps

import (
	"github.com/busy-cloud/boat/lib"
)

var _apps lib.Map[App]

func Register(a *App) {
	a.Internal = true
	_apps.Store(a.Id, a)
}

func Unregister(id string) {
	a := _apps.Load(id)
	if a != nil && a.Internal {
		_apps.Delete(id)
	}
}
