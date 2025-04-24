package table

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/config"
)

func init() {
	boot.Register("table", &boot.Task{
		Startup: Startup,
		Depends: []string{"config", "database"},
	})
}

func Startup() error {
	var err error

	//加载表
	paths := config.GetStringSlice(MODULE, "paths")
	if len(paths) == 0 {
		for _, path := range paths {
			err = Scan(path)
			if err != nil {
				return err
			}
		}
	}

	//同步
	if config.GetBool(MODULE, "sync") {
		var tbs []*Table
		tables.Range(func(name string, tab *Table) bool {
			tbs = append(tbs, tab)
			return true
		})
		if len(tbs) > 0 {
			err = Sync(tbs)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
