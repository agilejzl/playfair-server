package main

import (
	_ "playfair-server/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

