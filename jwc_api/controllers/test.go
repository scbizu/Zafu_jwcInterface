package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {

	this.Ctx.WriteString("hello")
}
