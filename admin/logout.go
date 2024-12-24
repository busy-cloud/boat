package admin

import (
	"github.com/busy-cloud/boat/curd"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	u := session.Get("user")
	if u == nil {
		curd.Fail(ctx, "未登录")
		return
	}

	//user := u.(int64)
	//_, _ = db.Engine.InsertOne(&types.UserEvent{UserId: user, ModEvent: types.ModEvent{Type: "退出"}})

	session.Clear()
	_ = session.Save()
	curd.OK(ctx, nil)
}
