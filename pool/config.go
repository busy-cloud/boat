package pool

import (
	"github.com/busy-cloud/boat/config"
)

const MODULE = "pool"

func init() {
	config.Register(MODULE, "size", 10000)
}
