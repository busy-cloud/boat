package table

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ApiDelete(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	id := strings.TrimLeft(ctx.Param("id"), "/")
	cnt, err := table.DeleteById(id)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, cnt)
}
