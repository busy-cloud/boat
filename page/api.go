package page

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
	"path"
)

func init() {
	//api.Register("GET", "/pages", pageGet)
	api.Register("GET", "/page/:page", pageGet)
	api.Register("GET", ":app/page/:page", appPageGet)
}

// @Summary 获取页面
// @Schemes
// @Description 获取页面
// @Tags plugin
// @Param page path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]Page] 返回插件信息
// @Router /page/:page [get]
func pageGet(ctx *gin.Context) {
	page := ctx.Param("page")
	ctx.FileFromFS(page+".json", &pages)
}

// @Summary 获取页面
// @Schemes
// @Description 获取页面
// @Tags plugin
// @Param page path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]Page] 返回插件信息
// @Router /page/:page [get]
func appPageGet(ctx *gin.Context) {
	app := ctx.Param("app")
	page := ctx.Param("page")
	file := path.Join(app, page)
	ctx.FileFromFS(file+".json", &pages)
}
