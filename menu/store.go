package menu

import (
	"embed"
	"encoding/json"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/store"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"strings"
)

var menuStore store.Store

func Dir(dir string, base string) {
	menuStore.Dir(dir, base)
}

func Zip(zip string, base string) {
	menuStore.Zip(zip, base)
}

func EmbedFS(fs embed.FS, base string) {
	menuStore.EmbedFS(fs, base)
}

func Load() error {
	//var ms []*Menu
	for _, item := range menuStore.Items {
		entries, err := item.ReadDir("")
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			ext := filepath.Ext(entry.Name())
			if ext == ".json" {
				buf, err := item.ReadFile(entry.Name())
				if err != nil {
					return err
				}
				var menu Menu
				err = json.Unmarshal(buf, &menu)
				if err != nil {
					log.Error(err)
					continue
				}

				Register(strings.TrimSuffix(entry.Name(), ext), &menu)
			} else if ext == ".yaml" {
				buf, err := item.ReadFile(entry.Name())
				if err != nil {
					return err
				}
				var menu Menu
				err = yaml.Unmarshal(buf, &menu)
				if err != nil {
					log.Error(err)
					continue
				}

				Register(strings.TrimSuffix(entry.Name(), ext), &menu)
			}
		}
	}
	return nil
}
