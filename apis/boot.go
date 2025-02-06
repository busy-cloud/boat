package apis

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/web"
)

func init() {
	boot.Register("apis", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"config", "web"},
	})
}

func Startup() error {

	web.Engine.POST("api/login", login)

	return nil
}
