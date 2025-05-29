package menu

import (
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/mqtt"
	"github.com/busy-cloud/boat/version"
)

type Menu struct {
	Name string `json:"name"`
	Icon string `json:"icon,omitempty"`
	//Domain     []string `json:"domain"` //域 admin project 或 dealer等
	Privileges []string `json:"privileges,omitempty"`
	Items      []*Item  `json:"items"`
	Index      int      `json:"index"`
}

type Item struct {
	Name string `json:"name,omitempty"`
	//Type       string         `json:"type,omitempty"` //route 路由, web 嵌入web, window 独立弹出
	Url        string         `json:"url,omitempty"`
	Query      map[string]any `json:"query,omitempty"`
	Privileges []string       `json:"privileges,omitempty"`
}

var menus lib.Map[Menu]

func Register(name string, menu *Menu) {
	if version.Name == "" || version.Name == "boat" {
		menus.Store(name, menu)
	} else {
		mqtt.Publish("boat/register/menu/"+name, menu)
	}
}

func Unregister(name string) {
	menus.Delete(name)
}
