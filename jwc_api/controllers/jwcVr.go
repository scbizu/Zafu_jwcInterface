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
	Tfunc := models.Getvrcode(client)
	_, T_CO := Tfunc()
	models.SetT_cookie(T_CO)
	j.TplName = "vrcode.tpl"
}
