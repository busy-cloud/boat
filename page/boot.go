package page

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/web"
	"github.com/gin-gonic/gin"
)

func init() {
	boot.Register("pages", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"config", "web"},
	})
}

func Startup() error {

	//页面
	web.Engine().GET("page/*page", func(ctx *gin.Context) {
		page := ctx.Param("page")
		ctx.FileFromFS(page+".json", &pages)
	})

	return nil
}
