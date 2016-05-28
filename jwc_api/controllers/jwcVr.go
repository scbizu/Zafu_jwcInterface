package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface.git/jwc_api/models"
)

type JwcVrController struct {
	beego.Controller
}

func (j *JwcVrController) Get() {
	client := models.InitClient()
	_ = models.Getvrcode(client)
	j.TplName = "vrcode.tpl"
}
