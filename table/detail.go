package table

import (
	"github.com/gin-gonic/gin"
)

func ApiDetail(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	doc, err := table.Get(ctx.Param("id"), nil)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, doc)
}
