package page

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {
	api.Register("GET", "page/*page", pageGet)
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
