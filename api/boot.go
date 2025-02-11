package api

import (
	"github.com/busy-cloud/boat/boot"
)

func init() {
	boot.Register("api", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: nil,
		Depends:  []string{"web", "log", "database", "apis"},
	})
}

func Startup() error {

	registerRoutes("api")

	return nil
}
