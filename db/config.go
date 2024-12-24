package db

import (
	"github.com/busy-cloud/boat/config"
)

const MODULE = "database"

func init() {
	config.Register(MODULE, "type", "sqlite3")
	config.Register(MODULE, "url", "boat.db") //"root:root@tcp(localhost:3306)/master?charset=utf8"
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "sync", true)
}

func SetDatabaseUrl(url string) {
	config.Set(MODULE, "url", url)
}
