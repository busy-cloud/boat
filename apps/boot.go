package apps

import (
	"encoding/json"
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/web"
	"go.uber.org/multierr"
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

	//注册
	pages.Dir("pages", "")

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

			//记录目录
			//p.dir = path.Join(dir, d.Name())

			_apps.Store(d.Name(), &p)
		}
	}

	return err
}
