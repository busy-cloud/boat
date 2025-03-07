package menu

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {
	//api.Register("GET", "/menu/:domain", menuGet)
	api.Register("GET", "/menus", menuGet)
}

// @Summary 获取菜单
// @Schemes
// @Description 获取菜单
// @Tags plugin
// @Param domain path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]Menu] 返回插件信息
// @Router /menus [get]
func menuGet(ctx *gin.Context) {
	//domain := ctx.Param("domain")

	//TODO 获取用户权限，过滤菜单
	var ms []*Menu
	menus.Range(func(name string, m *Menu) bool {
		ms = append(ms, m)
		return true
	})
	api.OK(ctx, ms)
}
