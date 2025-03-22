package apis

import "github.com/busy-cloud/boat/config"

const MODULE = "apis"

func init() {
	config.SetDefault(MODULE, "password", md5hash("123456"))
}
