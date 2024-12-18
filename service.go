package boat

import (
	"github.com/kardianos/service"
)

var config = service.Config{
	Name:        "boat",
	DisplayName: "Boat",
	Description: "Process Manager for General IoT Backend",
	//Arguments        []string
	//Option: service.KeyValue{
	//	"Restart":                "always",
	//	"OnFailure":              "restart",
	//	"OnFailureDelayDuration": "5s",
	//	"OnFailureResetPeriod":   30,
	//},
}
