package web

import (
	"github.com/busy-cloud/boat/log"
	"mime"
)

func init() {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		log.Error(err)
	}
}
