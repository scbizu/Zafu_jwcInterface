package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface/jwc_api/models"
)

type JwcVrController struct {
	beego.Controller
}

func (j *JwcVrController) Get() {
	client := models.InitClient()
	_ = models.Getvrcode(client)
	j.TplName = "vrcode.tpl"
	//	j.Redirect("http://210.33.60.8/CheckCode.aspx", 302)
}
