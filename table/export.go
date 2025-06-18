package table

import (
	"github.com/busy-cloud/boat/curd"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ApiExport(ctx *gin.Context) {
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

	//查询
	results, err := table.Find(body.Filter, []string{"*"}, int(body.Skip), int(body.Limit))
	if err != nil {
		Error(ctx, err)
		return
	}

	buf, err := sonic.Marshal(results)
	if err != nil {
		Error(ctx, err)
		return
	}

	filename := table.Name + "-export-" + time.Now().Format("20060102150405") + ".json"
	// 设置响应头
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Length", strconv.Itoa(len(buf)))
	_, _ = ctx.Writer.Write(buf)
	//ctx.Abort()
}
