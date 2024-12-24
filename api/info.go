package api

import (
	"github.com/busy-cloud/boat/build"
	"github.com/busy-cloud/boat/curd"
	"github.com/gin-gonic/gin"
	mochi "github.com/mochi-mqtt/server/v2"
	"runtime"
)

func info(ctx *gin.Context) {
	curd.OK(ctx, gin.H{
		"runtime": runtime.Version(),
		"build":   build.Build,
		"version": build.Version,
		"git":     build.GitHash,
		"gin":     gin.Version,
		"mqtt":    mochi.Version,
	})
}
