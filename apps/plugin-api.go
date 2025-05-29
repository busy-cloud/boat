package apps

import (
	_ "embed"
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/app"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

//go:embed icon.png
var defaultIcon []byte

const gmtFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

func init() {

	//启动时间作为默认图标的修改时间
	bootTime := time.Now()

	api.Register("GET", "plugin/list", func(ctx *gin.Context) {
		var as []*app.Plugin
		plugins.Range(func(name string, item *app.Plugin) bool {
			as = append(as, item)
			return true
		})
		api.OK(ctx, as)
	})

	api.Register("GET", "plugin/:app", func(ctx *gin.Context) {
		app := plugins.Load(ctx.Param("app"))
		if app != nil {
			api.Fail(ctx, "找不到插件")
			return
		}
		api.OK(ctx, app)
	})

	api.Register("GET", "plugin/:app/icon", func(ctx *gin.Context) {
		app := plugins.Load(ctx.Param("app"))
		if app != nil {
			api.Fail(ctx, "找不到插件")
			return
		}

		icon := app.Icon
		if icon == "" {
			icon = IconName
		}

		file, err := os.Open(IconName)
		if err != nil {
			//return nil, err
			//return icon, nil //使用默认图片
			ctx.Header("Last-Modified", bootTime.UTC().Format(gmtFormat))
			ctx.Header("Content-Type", "image/png")
			_, _ = ctx.Writer.Write(defaultIcon)
			return
		}
		defer file.Close()

		st, _ := file.Stat()
		buf, err := io.ReadAll(file)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.Header("Last-Modified", st.ModTime().UTC().Format(gmtFormat))
		ctx.Header("Content-Type", "image/png")
		_, _ = ctx.Writer.Write(buf)
	})

	api.RegisterAdmin("POST", "plugin/import", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			api.Error(ctx, err)
			return
		}
		f, err := file.Open()
		if err != nil {
			api.Error(ctx, err)
			return
		}
		defer f.Close()

		f2, err := os.CreateTemp("plugin", "install-*")
		if err != nil {
			api.Error(ctx, err)
			return
		}
		defer f2.Close()

		_, err = io.Copy(f2, f)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		err = app.Unpack(app.PublicKey(), f2.Name(), RootPath)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		api.OK(ctx, nil)
	})

	api.RegisterAdmin("GET", "plugin/:app/install", func(ctx *gin.Context) {

	})

	api.RegisterAdmin("GET", "plugin/:app/delete", func(ctx *gin.Context) {

	})

}
