package menu

import (
	"encoding/json"
	"github.com/busy-cloud/boat/app"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/mqtt"
	"os"
	"path/filepath"
	"strings"
)

type Menu struct {
	Name string `json:"name"`
	Icon string `json:"icon,omitempty"`
	//Domain     []string `json:"domain"` //域 admin project 或 dealer等
	Privileges []string `json:"privileges,omitempty"`
	Items      []*Item  `json:"items"`
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
	if app.Name == "" {
		menus.Store(name, menu)
	} else {
		mqtt.Publish("boat/register/menu/"+name, menu)
	}
}

func Unregister(name string) {
	menus.Delete(name)
}

func Load(dir string) {
	_ = os.MkdirAll(dir, os.ModePerm)
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Error(err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		if ext == ".json" {
			fn := filepath.Join(dir, entry.Name())
			buf, err := os.ReadFile(fn)
			if err != nil {
				log.Error(err)
				continue
			}
			var menu Menu
			err = json.Unmarshal(buf, &menu)
			if err != nil {
				log.Error(err)
				continue
			}

			name := strings.TrimSuffix(entry.Name(), ext)
			Register(name, &menu)
		}
	}
}
