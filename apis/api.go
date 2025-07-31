package apis

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {
	//api.RegisterUnAuthorized("POST", "login", login)
	//api.Register("GET", "logout", logout)

	//api.Register("GET", "me", me)

	//api.RegisterAdmin("POST", "password", password)

}

func me(ctx *gin.Context) {
	id := ctx.GetString("user")

	if id == "" {
		api.Fail(ctx, "未登录")
		return
	}

	api.OK(ctx, gin.H{"id": id})
}
