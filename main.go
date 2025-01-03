package main

import (
	_ "github.com/busy-cloud/boat/admin"
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/boat/broker"
	_ "github.com/busy-cloud/boat/internal"
	"github.com/busy-cloud/boat/log"
	_ "github.com/busy-cloud/boat/mqtt"
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

	//注册接口
	api.RegisterRoutes(web.Engine.Group("api"))

	//启动服务
	err = web.Serve()
	if err != nil {
		log.Error(err)
	}
}
