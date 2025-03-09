package plugin

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/web"
)

func init() {
	boot.Register("plugin", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"web"},
	})
}

func Startup() error {
	web.Engine().Use(Proxy)
	return nil
}
