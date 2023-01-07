package main

import (
	_ "github.com/mattn/go-sqlite3"

	_ "cppm/routers"
	beego "github.com/beego/beego/v2/server/web"
	_ "cppm/models"
)

func init() {
}

func main() {
	beego.Run()
}

