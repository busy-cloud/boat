package table

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
	"slices"
)

func ApiCount(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var body ParamSearch
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	//多租户过滤
	tid := ctx.GetString("tid")
	if tid != "" {
		tenantId := slices.IndexFunc(table.Fields, func(field *Field) bool {
			return field.Name == "tenant_id"
		})
		if tenantId > -1 {
			//只有未传值tenant_id时，才会赋值用户所在的tenant_id
			if _, ok := body.Filter["tenant_id"]; !ok {
				body.Filter["tenant_id"] = tid
			}
		}
	}

	ret, err := table.Count(body.Filter)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, ret)
}
