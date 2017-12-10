package main

import (
	"github.com/astaxie/beego"
	_ "github.com/gabezeck/playlistr/routers"
)

func main() {
	beego.Run()
}
