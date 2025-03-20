package menu

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/mqtt"
	"strings"
)

func init() {
	boot.Register("menu", &boot.Task{
		Startup: Startup,
		Depends: []string{"web", "mqtt"},
	})
}

func Startup() error {

	mqtt.SubscribeStruct[Menu]("boat/register/menu/+", func(topic string, menu *Menu) {
		name := strings.TrimPrefix(topic, "boat/register/menu/")
		menus.Store(name, menu)
	})

	return Load()
}
