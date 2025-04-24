package table

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var tables []*Table

func Register(table *Table) {
	tables = append(tables, table)
}

func Load(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".json" {
			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var table Table
			err = json.Unmarshal(buf, &table)
			if err != nil {
				return err
			}

			tables = append(tables, &table)
		}
		return nil
	})
}
