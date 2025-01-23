package admin

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {
	api.Register("GET", "me", me)
	api.Register("GET", "logout", logout)
	api.Register("POST", "password", password)
}

func me(ctx *gin.Context) {
	id := ctx.GetString("user")

	if id == "" {
		api.Fail(ctx, "未登录")
		return
	}

	api.OK(ctx, gin.H{"id": id})
}
