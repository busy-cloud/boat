package main

import (
	_ "github.com/busy-cloud/boat/apis"
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/boat/broker"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/plugin"
	"github.com/busy-cloud/boat/web"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	viper.SetConfigName("boat")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs

		//关闭web，出发
		_ = web.Shutdown()
	}()

	//安全退出
	defer boot.Shutdown()

	err := boot.Startup()
	if err != nil {
		log.Error(err)
		return
	}

	//执行插件代理
	web.Engine.Use(plugin.Proxy)

	//启动服务
	err = web.Serve()
	if err != nil {
		log.Error(err)
	}
}
