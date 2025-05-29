package apis

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/version"
	"github.com/gin-gonic/gin"
	mochi "github.com/mochi-mqtt/server/v2"
	"runtime"
)

func info(ctx *gin.Context) {
	api.OK(ctx, gin.H{
		"runtime": runtime.Version(),
		"build":   version.Build,
		"version": version.Version,
		"git":     version.GitHash,
		"gin":     gin.Version,
		"broker":  mochi.Version,
	})
}
