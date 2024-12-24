package service

import (
	"github.com/busy-cloud/boat/config"
)

const MODULE = "service"

func init() {
	config.Register(MODULE, "name", "boat")
	config.Register(MODULE, "display", "Boat")
	config.Register(MODULE, "description", "Process Manager for General IoT Backend")
	config.Register(MODULE, "arguments", []string{})
}
