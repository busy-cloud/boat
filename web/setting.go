package web

import (
	"github.com/busy-cloud/boat/config"
	"github.com/busy-cloud/boat/smart"
)

func init() {
	config.Register(MODULE, &config.Form{
		Title:  "Web配置",
		Module: MODULE,
		Form: smart.Form{
			Fields: []smart.Field{
				{Key: "port", Label: "端口", Type: "number", Required: true, Default: 8080, Min: 1, Max: 65535},
				{Key: "debug", Label: "调试模式", Type: "switch"},
				{Key: "cors", Label: "跨域请求", Type: "switch"},
				{Key: "gzip", Label: "压缩模式", Type: "switch"},
				{
					Key: "https", Label: "HTTPS", Type: "select",
					Options: []smart.SelectOption{
						{Label: "禁用", Value: ""},
						{Label: "TLS", Value: "TLS"},
						{Label: "LetsEncrypt", Value: "LetsEncrypt"},
					},
				},
				{Key: "tls_cert", Label: "证书cert", Type: "file"},
				{Key: "Key", Label: "证书key", Type: "file"},
				{Key: "email", Label: "E-Mail", Type: "text"},
				{Key: "hosts", Label: "域名", Type: "tags", Default: []string{}},
			},
		},
	})
}
