package plugin

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
	boot.Register("plugin", &boot.Task{
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

	plugins.Range(func(name string, plugin *Plugin) bool {
		//err = multierr.Append(err, plugin.Open())
		err := plugin.Open()
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
	plugins.Range(func(name string, plugin *Plugin) bool {
		err = multierr.Append(err, plugin.Close())
		return true
	})
	return
}

func load() error {
	dir := path.Join("plugins")
	_ = os.MkdirAll(dir, os.ModePerm)

	ds, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	//加载
	for _, d := range ds {
		if d.IsDir() {
			pp := path.Join(dir, d.Name(), "plugin.json")
			buf, e := os.ReadFile(pp)
			if e != nil {
				err = multierr.Append(err, e)
				continue
			}

			var plugin Plugin
			e = json.Unmarshal(buf, &plugin)
			if e != nil {
				err = multierr.Append(err, e)
				continue
			}

			//记录目录
			plugin.dir = path.Join(dir, d.Name())

			plugins.Store(d.Name(), &plugin)
		}
	}

	return err
}
