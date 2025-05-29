package apps

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Proxy(ctx *gin.Context) {

	if str, has := strings.CutPrefix(ctx.Request.RequestURI, "/plugin/"); has {
		if len(str) == 0 {
			return
		}
		if app, _, has := strings.Cut(str, "/"); has {
			if p := plugins.Load(app); p != nil {
				p.ServeApi(ctx)
				ctx.Abort()
				return
			}
		}
	}

	ctx.Next()
}
