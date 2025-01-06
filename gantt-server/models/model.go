package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func NewEngine(addr string) error {
	var err error
	engine, err = xorm.NewEngine("sqlite3", addr)
	if err != nil {
		return fmt.Errorf("xorm NewEngine failed, error: %s", err.Error())
	}
	err = syncModels()
	return nil
}

func syncModels() error {
	err := engine.Sync(new(Task))
	if err != nil {
		return fmt.Errorf("xorm sync failed, error: %s", err.Error())
	}
	return nil
}
