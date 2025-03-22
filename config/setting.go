package config

import (
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/smart"
)

type Form struct {
	smart.Form

	Title  string `json:"title"`
	Module string `json:"module"`
}

var modules lib.Map[Form]

func Register(module string, form *Form) {
	modules.Store(module, form)
}

func Unregister(module string) {
	modules.Delete(module)
}

func GetModule(module string) *Form {
	return modules.Load(module)
}

func GetModules() []Form {
	var ms []Form
	modules.Range(func(_ string, item *Form) bool {
		m := *item
		m.Form.Fields = nil
		ms = append(ms, m)
		return true
	})
	return ms
}
