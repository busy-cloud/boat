package apis

import "github.com/busy-cloud/boat/config"

const MODULE = "apis"

func init() {
	config.Register(MODULE, "password", md5hash("123456"))
}
