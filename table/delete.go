package table

import (
	"github.com/gin-gonic/gin"
)

func ApiDelete(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	//ids := ctx.QueryArray("id") //依次删除

	cnt, err := table.Delete(Document{"id": ctx.Param("id")})
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, cnt)
}
