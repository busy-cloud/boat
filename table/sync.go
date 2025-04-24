package table

import (
	"context"
	"github.com/busy-cloud/boat/db"
	"xorm.io/xorm/schemas"
)

func Sync(tables []*Table) error {

	tbs, err := db.Engine().Dialect().GetTables(db.Engine().DB(), context.Background())
	if err != nil {
		return err
	}

	for _, table := range tables {

		var t *schemas.Table
		for _, tb := range tbs {
			if tb.Name == table.Name {
				t = tb
				break
			}
		}

		//新表
		if t == nil {
			err = table.Create()
			if err != nil {
				return err
			}
			continue
		}

		//更新列
		schema := table.Schema()
		for _, col := range schema.Columns() {
			found := false
			for _, c2 := range t.Columns() {
				if col.Name == c2.Name {
					found = true
					break
				}
			}

			//添加新列
			if !found {
				t.AddColumn(col)
				sql := db.Engine().Dialect().AddColumnSQL(t.Name, col)
				_, err = db.Engine().Exec(sql)
				if err != nil {
					return err
				}
			}

			//TODO 处理修改字段类型
		}

		//更新索引
		for _, index := range schema.Indexes {
			found := false
			for _, i := range t.Indexes {
				if index.Equal(i) {
					found = true
					break
				}
			}

			//添加索引
			if !found {
				sql := db.Engine().Dialect().CreateIndexSQL(t.Name, index)
				_, err = db.Engine().Exec(sql)
				if err != nil {
					return err
				}
			}

			//TODO 删除索引
		}

	}

	return err
}
