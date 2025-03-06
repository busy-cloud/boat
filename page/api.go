package page

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {
	//api.Register("GET", "/pages", menuGet)
	api.Register("GET", "/page/:page", menuGet)
}

// @Summary 获取页面
// @Schemes
// @Description 获取页面
// @Tags plugin
// @Param domain path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]Page] 返回插件信息
// @Router /page/:page [get]
func menuGet(ctx *gin.Context) {
	m := ctx.Param("page")
	md := pages.Load(m)
	if md == nil {
		api.Fail(ctx, "页面不存在")
		return
	}
	api.OK(ctx, md)
}
