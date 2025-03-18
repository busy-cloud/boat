package menu

import (
	"embed"
	"encoding/json"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/store"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

var menus2 store.Store

func Dir(dir string, base string) {
	menus2.Dir(dir, base)
}

func Zip(zip string, base string) {
	menus2.Zip(zip, base)
}

func EmbedFS(fs embed.FS, base string) {
	menus2.EmbedFS(fs, base)
}

func Load() ([]*Menu, error) {
	var ms []*Menu
	for _, item := range menus2.Items {
		entries, err := item.ReadDir("")
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			ext := filepath.Ext(entry.Name())
			if ext == ".json" {
				buf, err := item.ReadFile(entry.Name())
				if err != nil {
					return nil, err
				}
				var menu Menu
				err = json.Unmarshal(buf, &menu)
				if err != nil {
					log.Error(err)
					continue
				}
				ms = append(ms, &menu)
			} else if ext == ".yaml" {
				buf, err := item.ReadFile(entry.Name())
				if err != nil {
					return nil, err
				}
				var menu Menu
				err = yaml.Unmarshal(buf, &menu)
				if err != nil {
					log.Error(err)
					continue
				}
				ms = append(ms, &menu)
			}
		}
	}
	return ms, nil
}
