package test

import (
	"github.com/busy-cloud/boat/db"
	"testing"
)

type Abc struct {
	Id string `json:"id" xorm:"pk"`
}

func init() {
	db.Register(&Abc{})
}

func TestDatabaseGet(t *testing.T) {
	_ = db.Startup()
	var a Abc
	has, err := db.Engine.ID("123").Get(&a)
	t.Log(has, err)
}
