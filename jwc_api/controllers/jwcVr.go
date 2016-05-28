package controllers

import (
	"os/exec"

	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface/jwc_api/models"
)

type JwcVrController struct {
	beego.Controller
}

func (j *JwcVrController) Get() {
	client := models.InitClient()
	_ = models.Getvrcode(client)
	Mvcmd := exec.Command("mv", "verify.gif", "code")
	Mvcmd.Run()
	j.TplName = "vrcode.tpl"
}
