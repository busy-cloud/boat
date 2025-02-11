package apis

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginObj struct {
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func md5hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	var obj loginObj
	if err := ctx.ShouldBind(&obj); err != nil {
		api.Error(ctx, err)
		return
	}

	password := config.GetString(MODULE, "password")
	if password != obj.Password {
		api.Fail(ctx, "密码错误")
		return
	}

	//_, _ = db.Engine.InsertOne(&types.UserEvent{UserId: user.id, ModEvent: types.ModEvent{Type: "登录"}})

	//存入session
	session.Set("user", "admin")
	session.Set("admin", true)
	_ = session.Save()

	api.OK(ctx, gin.H{"id": "admin", "admin": true})
}
