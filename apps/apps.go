package apps

import (
	"github.com/busy-cloud/boat/app"
	"github.com/busy-cloud/boat/lib"
)

var _apps lib.Map[App]

func Register(a *app.App) {
	aa := &App{App: *a}
	aa.Internal = true
	_apps.Store(a.Id, aa)
}

func Unregister(id string) {
	a := _apps.Load(id)
	if a != nil && a.Internal {
		_apps.Delete(id)
	}
}
