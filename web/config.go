package web

import (
	"github.com/busy-cloud/boat/config"
	"math/rand/v2"
)

const MODULE = "web"

func init() {
	config.Register(MODULE, "mode", "http")               //http, https, tls, ssl, autocert, letsencrypt, unix
	config.Register(MODULE, "port", 8000+rand.UintN(100)) //端口号 8000 - 8099
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "cors", false)
	config.Register(MODULE, "gzip", true)
	config.Register(MODULE, "tls_cert", "")
	config.Register(MODULE, "tls_key", "")
	config.Register(MODULE, "hosts", []string{}) //域名
	config.Register(MODULE, "email", "")
	config.Register(MODULE, "jwt_key", "boat")
	config.Register(MODULE, "jwt_expire", 24*30) //小时
	config.Register(MODULE, "unix_socket", "")

	//通过环境变量，强制修改监听类型和端口
	//BOAT_WEB.MODE=unix;
	//BOAT_WEB.UNIX_SOCKET=boat.sock

}
