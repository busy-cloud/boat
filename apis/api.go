package apis

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {

	api.RegisterUnAuthorized("GET", "oem", oem)
	api.RegisterAdmin("POST", "oem/update", oemUpdate)
	api.RegisterUnAuthorized("GET", "info", info)
	api.RegisterUnAuthorized("POST", "login", login)
	api.Register("GET", "me", me)
	api.Register("GET", "logout", logout)
	api.RegisterAdmin("POST", "password", password)

	//router.GET("/oem", apis2.oem)
	//router.GET("/info", apis2.info)
	//apis2.oemRouter(router.Group("/oem"))
}

func me(ctx *gin.Context) {
	id := ctx.GetString("user")

	if id == "" {
		api.Fail(ctx, "未登录")
		return
	}

	api.OK(ctx, gin.H{"id": id})
}
