package broker

import (
	"github.com/god-jason/boat/boot"
	"github.com/god-jason/boat/config"
	"github.com/god-jason/boat/web"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"log/slog"
	"os"
)

func init() {
	boot.Register("broker", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "log", "database"},
	})
}

var server *mqtt.Server

func Startup() (err error) {
	//禁用不启动
	if !config.GetBool(MODULE, "enable") {
		return nil
	}

	opts := &mqtt.Options{
		InlineClient: true,
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		})),
	}
	server = mqtt.New(opts)

	if config.GetBool(MODULE, "anonymous") {
		err = server.AddHook(new(auth.AllowHook), nil)
	} else {
		err = server.AddHook(new(Hook), nil)
	}

	if err != nil {
		return err
	}

	//内置监听
	err = server.AddListener(listeners.NewTCP(listeners.Config{
		ID:      "base",
		Address: ":" + config.GetString(MODULE, "port"),
	}))
	if err != nil {
		return err
	}

	//启用unixsock，速度更快
	err = server.AddListener(listeners.NewUnixSock(listeners.Config{
		ID:      "unix",
		Address: config.GetString(MODULE, "unixsock"),
	}))
	if err != nil {
		return err
	}

	//监听Websocket
	web.Engine.GET("/mqtt", GinBridge)

	return server.Serve()
}

func Shutdown() error {
	if server != nil {
		return server.Close()
	}
	return nil
}
