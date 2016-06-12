package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface/jwc_api/models"
)

type JwcVrController struct {
	beego.Controller
}

func (j *JwcVrController) Get() {
	models.StuNo = j.GetString("stuno")
	beego.Debug(models.StuNo)
	client := models.InitClient()
	beego.Debug(client)
	c, T_CO := models.Getvrcode(client)
	models.SetT_cookie(T_CO)
	beego.Debug(T_CO)
	models.Client = c
	//	beego.Debug(c)
	j.Data["StuNo"] = models.StuNo
	j.TplName = "vrcode.tpl"
}
