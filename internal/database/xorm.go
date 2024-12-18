package database

import (
	_ "github.com/go-sql-driver/mysql" //mysql
	_ "github.com/jackc/pgx/v5"        //postgresql pgx/v5
	_ "github.com/mattn/go-sqlite3"    //sqlite3
	"xorm.io/xorm"
)

var engine *xorm.Engine

func Start() (err error) {
	engine, err = xorm.NewEngine("sqlite3", "boat.db")
	return
}
