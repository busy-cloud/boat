package db

import (
	"github.com/god-jason/boat/setting"
	"github.com/god-jason/boat/smart"
)

func init() {
	setting.Register(MODULE, &setting.Module{
		Name:   "数据库",
		Module: MODULE,
		Title:  "数据库配置",
		Form: []smart.Field{
			{
				Key: "Type", Label: "数据库类型", Type: "select", Default: "sqlite3",
				Options: []smart.SelectOption{
					{Label: "SQLite（内置）", Value: "sqlite3"},
					{Label: "MySQL", Value: "mysql"},
					{Label: "Postgres SQL", Value: "postgres"},
					//{Label: "MS SQL Server", Value: "sqlserver"},
					//{Label: "Oracle", Value: "godror"},
				},
			},
			{Key: "url", Label: "连接字符串", Type: "text"},
			{Key: "debug", Label: "调试模式", Type: "switch"},
		},
	})
}
