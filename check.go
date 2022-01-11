package main

import (
	"database/sql"
	"github.com/beego/beego/v2/adapter/toolbox"
)

type DatabaseCheck struct {
}

func (d *DatabaseCheck) Check() error {
	_, err := sql.Open("mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/forum?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
}
