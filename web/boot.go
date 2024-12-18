package web

import "github.com/god-jason/boat/boot"

func init() {
	boot.Register("web", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"config"},
	})
}
