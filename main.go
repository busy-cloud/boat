package main

import (
	_ "github.com/busy-cloud/boat/apis"
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/boat/broker"
	"github.com/busy-cloud/boat/log"
	_ "github.com/busy-cloud/boat/menu"
	"github.com/busy-cloud/boat/page"
	_ "github.com/busy-cloud/boat/page"
	"github.com/busy-cloud/boat/plugin"
	"github.com/busy-cloud/boat/service"
	_ "github.com/busy-cloud/boat/setting"
	"github.com/busy-cloud/boat/web"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	//模板页面
	page.Dir("pages", "")

	//异步执行，避免堵塞
	go func() {
		//启动服务
		err := web.Serve()
		if err != nil {
			//安全退出
			//_ = boot.Shutdown()
			log.Error(err)
		}
	}()

	log.Info("main started")

	return nil
}

func Shutdown() error {
	log.Info("main shutdown")

	return boot.Shutdown()
}

func main() {
	help := pflag.BoolP("help", "h", false, "show help")
	install := pflag.BoolP("install", "i", false, "install as service")
	uninstall := pflag.BoolP("uninstall", "u", false, "uninstall service")

	pflag.Parse()
	if *help {
		pflag.PrintDefaults()
		return
	}

	err := service.Register(Startup, Shutdown)
	if err != nil {
		log.Fatal(err)
	}

	if *install {
		log.Info("install service")
		err = service.Install()
		if err != nil {
			log.Fatal(err)
		}
	} else if *uninstall {
		log.Info("uninstall service")
		err = service.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigs
		log.Info("signal received ", s)

		//_ = boot.Shutdown()
		err := service.Stop()
		if err != nil {
			log.Error(err)
		}

		time.AfterFunc(10*time.Second, func() {
			os.Exit(0)
		})
	}()

	err = service.Run()
	if err != nil {
		log.Fatal(err)
	}

	println("bye")
}
