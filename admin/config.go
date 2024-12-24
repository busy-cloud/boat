package admin

import "github.com/busy-cloud/boat/config"

const MODULE = "admin"

func init() {
	config.Register(MODULE, "password", md5hash("123456"))
}
