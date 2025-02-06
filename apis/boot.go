package apis

import (
	"github.com/busy-cloud/boat/boot"
)

func init() {
	boot.Register("apis", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"config", "web"},
	})
}

func Startup() error {

	return nil
}
