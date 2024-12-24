package broker

import (
	"github.com/busy-cloud/boat/setting"
	"github.com/busy-cloud/boat/smart"
)

func init() {
	setting.Register(MODULE, &setting.Module{
		Name:   "MQTT总线",
		Module: MODULE,
		Title:  "MQTT总线配置",
		Form: []smart.Field{
			{Key: "enable", Label: "启用", Type: "switch", Required: true},
			{Key: "anonymous", Label: "支持匿名访问", Type: "switch"},
			{Key: "port", Label: "端口", Type: "number", Required: true, Default: 1883},
			{Key: "unixsock", Label: "UnixSock", Type: "text"},
		},
	})
}
