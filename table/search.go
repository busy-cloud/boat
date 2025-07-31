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

	//多租户过滤
	tid := ctx.GetString("tid")
	if tid != "" {
		field := table.Field("tenant_id")
		if field != nil {
			//只有未传值tenant_id时，才会赋值用户所在的tenant_id
			if _, ok := body.Filter["tenant_id"]; !ok {
				body.Filter["tenant_id"] = tid
			}
		}
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
