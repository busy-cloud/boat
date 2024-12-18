package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/boat/build"
	"github.com/god-jason/boat/curd"
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
