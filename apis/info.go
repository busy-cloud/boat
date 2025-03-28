package apis

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/app"
	"github.com/gin-gonic/gin"
	mochi "github.com/mochi-mqtt/server/v2"
	"runtime"
)

func info(ctx *gin.Context) {
	api.OK(ctx, gin.H{
		"runtime": runtime.Version(),
		"build":   app.Build,
		"version": app.Version,
		"git":     app.GitHash,
		"gin":     gin.Version,
		"broker":  mochi.Version,
	})
}
