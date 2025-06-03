package apps

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

var pages store.Store

func Pages() *store.Store {
	return &pages
}

func init() {
	api.Register("GET", "page/:app/*page", func(ctx *gin.Context) {
		k := ctx.Param("app")

		//TODO先查询pages目录???

		//从应用列表中获取
		a := _apps.Load(k)
		if a != nil {
			if a.PagesFS != nil {
				ctx.FileFromFS(ctx.Param("page")+".json", http.FS(a.PagesFS)) //TODO 每次都创建了
			} else {
				ctx.String(http.StatusNotFound, "page not found")
			}
			return
		}

		ctx.String(http.StatusNotFound, "app not found")
	})
}
