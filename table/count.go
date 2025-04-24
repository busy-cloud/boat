package table

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/gin-gonic/gin"
)

func ApiCount(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var body curd.ParamSearch
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	ret, err := table.Count(body.Filter)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, ret)
}
