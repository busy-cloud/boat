package broker

import (
	"github.com/busy-cloud/boat/config"
	"os"
	"runtime"
)

const MODULE = "broker"

func init() {
	config.Register(MODULE, "enable", true)
	config.Register(MODULE, "anonymous", false)
	config.Register(MODULE, "port", 1883)

	if runtime.GOOS == "windows" {
		config.Register(MODULE, "unixsock", os.TempDir()+"/boat.sock")
	} else {
		config.Register(MODULE, "unixsock", "/var/run/boat.sock")
	}

	config.Register(MODULE, "loglevel", "ERROR")
}
