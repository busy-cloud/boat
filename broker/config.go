package broker

import (
	"github.com/busy-cloud/boat/config"
	"os"
)

const MODULE = "broker"

func init() {
	config.Register(MODULE, "enable", true)
	config.Register(MODULE, "anonymous", false)
	config.Register(MODULE, "port", 1883)
	config.Register(MODULE, "unixsock", os.TempDir()+"/boat.sock")
}
