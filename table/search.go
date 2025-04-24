package table

import (
	"github.com/busy-cloud/boat/curd"
	"github.com/gin-gonic/gin"
)

func ApiSearch(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}
	var body curd.ParamSearch
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		Error(ctx, err)
		return
	}

	cnt, err := table.Count(body.Filter)
	if err != nil {
		Error(ctx, err)
		return
	}

	results, err := table.Find(body.Filter, []string{"*"}, body.Skip, body.Limit)
	if err != nil {
		Error(ctx, err)
		return
	}

	//OK(ctx, results)
	List(ctx, results, cnt)
}
