package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "regInfo/routers"
)

func main() {
	beego.Run()
}
