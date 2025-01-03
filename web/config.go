package web

import (
	"github.com/busy-cloud/boat/config"
	"github.com/busy-cloud/boat/lib"
)

const MODULE = "web"

func init() {
	config.Register(MODULE, "port", 8008)
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "cors", false)
	config.Register(MODULE, "gzip", true)
	config.Register(MODULE, "https", "")
	config.Register(MODULE, "cert", "")
	config.Register(MODULE, "key", "")
	config.Register(MODULE, "hosts", []string{}) //域名
	config.Register(MODULE, "email", "")
	config.Register(MODULE, "id", "xid")
	config.Register(MODULE, "jwt_key", lib.RandomString(8))
	config.Register(MODULE, "jwt_expire", 24*30) //小时

}
