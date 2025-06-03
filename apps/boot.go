package apps

import (
	"encoding/json"
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/web"
	"github.com/gin-gonic/gin"
	"go.uber.org/multierr"
	"net/http"
	"os"
	"path"
)

func init() {
	boot.Register("apps", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"web"},
	})
}

func Startup() error {
	err := load()
	if err != nil {
		return err
	}

	_apps.Range(func(name string, p *App) bool {
		if len(p.Dependencies) > 0 {
			for _, d := range p.Dependencies {
				pp := _apps.Load(d)
				if pp == nil {
					err := pp.Open()
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
		//TODO 循环依赖问题

		//err = multierr.Append(err, internal.Open())
		err := p.Open()
		if err != nil {
			log.Error(err)
		}
		return true
	})

	//应用资源
	web.Engine().GET("assets/:app/*asset", func(ctx *gin.Context) {
		k := ctx.Param("app")

		//从应用列表中获取
		a := _apps.Load(k)
		if a != nil {
			if a.AssetsFS != nil {
				ctx.FileFromFS(ctx.Param("asset"), http.FS(a.AssetsFS)) //TODO 每次都创建了
			} else {
				ctx.String(http.StatusNotFound, "asset not found")
			}
			return
		}

		ctx.String(http.StatusNotFound, "app not found")
		//默认目录
		//ctx.FileFromFS(ctx.Param("asset"), http.Dir("assets"))
	})

	//注册到web引擎上
	web.Engine().Use(Proxy)

	return err
}

func Shutdown() (err error) {
	_apps.Range(func(name string, plugin *App) bool {
		err = multierr.Append(err, plugin.Close())
		return true
	})
	return
}

func load() error {
	dir := path.Join(RootPath)
	_ = os.MkdirAll(dir, os.ModePerm)

	ds, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	//加载
	for _, d := range ds {
		if d.IsDir() {
			pp := path.Join(dir, d.Name(), ManifestName)
			buf, e := os.ReadFile(pp)
			if e != nil {
				err = multierr.Append(err, e)
				continue
			}

			var p App
			e = json.Unmarshal(buf, &p)
			if e != nil {
				err = multierr.Append(err, e)
				continue
			}

			//强制修改为外部插件
			p.Internal = false

			//记录目录
			//p.dir = path.Join(dir, d.Name())

			_apps.Store(d.Name(), &p)
		}
	}

	return err
}
