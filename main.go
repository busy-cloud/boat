package main

import (
	_ "github.com/busy-cloud/boat/apis"
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/boat/broker"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/plugin"
	"github.com/busy-cloud/boat/service"
	"github.com/busy-cloud/boat/web"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func Startup() error {
	viper.SetConfigName("boat")

	err := boot.Startup()
	if err != nil {
		//_ = boot.Shutdown()
		return err
	}

	//执行插件代理
	web.Engine.Use(plugin.Proxy)

	//异步执行，避免堵塞
	go func() {
		//启动服务
		err := web.Serve()
		if err != nil {
			//安全退出
			_ = boot.Shutdown()
			log.Error(err)
		}
	}()

	return nil
}

func Shutdown() error {
	return boot.Shutdown()
}

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs

		_ = boot.Shutdown()
	}()

	err := service.Register(Startup, Shutdown)
	if err != nil {
		log.Fatal(err)
	}

	err = service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
