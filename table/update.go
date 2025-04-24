package table

import (
	"github.com/gin-gonic/gin"
)

func ApiUpdate(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var update Document
	err = ctx.ShouldBindJSON(&update)
	if err != nil {
		Error(ctx, err)
		return
	}

	cnt, err := table.UpdateById(ctx.Param("id"), update)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, cnt)
}
