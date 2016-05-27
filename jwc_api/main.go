package main

import (
	"github.com/astaxie/beego"
	_ "github.com/scbizu/Zafu_jwcInterface/jwc_api/docs"
	_ "github.com/scbizu/Zafu_jwcInterface/jwc_api/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
