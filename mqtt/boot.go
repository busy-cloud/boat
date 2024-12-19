package mqtt

import "github.com/god-jason/boat/boot"

func init() {
	boot.Register("mqtt", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"config", "broker"},
	})
}
