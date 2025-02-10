package plugin

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Proxy(ctx *gin.Context) {

	//插件 前端页面
	if str, has := strings.CutPrefix(ctx.Request.RequestURI, "/app/"); has {
		if app, _, has := strings.Cut(str, "/"); has {
			if p := plugins.Load(app); p != nil {
				if p.proxy == nil {
					//执行反向代理
					p.proxy.ServeHTTP(ctx.Writer, ctx.Request)
				}
			}
		}
		return
	}

	//插件 接口
	if str, has := strings.CutPrefix(ctx.Request.RequestURI, "/api/"); has {
		if app, _, has := strings.Cut(str, "/"); has {
			if p := plugins.Load(app); p != nil {
				if p.proxy == nil {
					//执行反向代理
					p.proxy.ServeHTTP(ctx.Writer, ctx.Request)
				}

			}
		}
		return
	}

}
