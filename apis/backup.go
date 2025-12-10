package apis

import (
	"bufio"
	"time"

	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
)

func init() {
	api.Register("GET", "backup", backupGet)
	api.Register("POST", "recover", recoverPost)
}

// 备份数据库
func backupGet(ctx *gin.Context) {
	engine := db.Engine()
	if engine == nil {
		api.Fail(ctx, "数据库未连接")
		return
	}

	// 设置响应头
	filename := "backup_" + time.Now().Format("20060102150405") + ".sql"
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "application/sql")

	writer := bufio.NewWriter(ctx.Writer)
	if err := engine.DumpAll(writer); err != nil {
		api.Error(ctx, err)
		return
	}
	writer.Flush()
}

// 恢复数据库
func recoverPost(ctx *gin.Context) {
	engine := db.Engine()
	if engine == nil {
		api.Fail(ctx, "数据库未连接")
		return
	}

	// 获取上传文件
	file, err := ctx.FormFile("file")
	if err != nil {
		api.Error(ctx, err)
		return
	}

	// 打开文件
	src, err := file.Open()
	if err != nil {
		api.Error(ctx, err)
		return
	}
	defer src.Close()

	// 恢复数据
	reader := bufio.NewReader(src)
	_, err = engine.Import(reader)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, "恢复成功")
}
