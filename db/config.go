package db

import (
	"github.com/busy-cloud/boat/config"
)

const MODULE = "database"

func init() {
	config.Register(MODULE, "type", "mysql")
	config.Register(MODULE, "url", "root:123456@tcp(localhost:3306)/boat?charset=utf8")
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "sync", true)
}

func SetDatabaseUrl(url string) {
	config.Set(MODULE, "url", url)
}
