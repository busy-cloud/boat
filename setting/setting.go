package setting

import (
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/smart"
)

type Form struct {
	Name   string     `json:"name"`
	Module string     `json:"module"`
	Title  string     `json:"title,omitempty"`
	Form   smart.Form `json:"-"`
}

var modules lib.Map[Form]

func Register(module string, form *Form) {
	modules.Store(module, form)
}

func Unregister(module string) {
	modules.Delete(module)
}

func Load(module string) *Form {
	return modules.Load(module)
}

func Modules() []*Form {
	var ms []*Form
	modules.Range(func(_ string, item *Form) bool {
		ms = append(ms, item)
		return true
	})
	return ms
}
