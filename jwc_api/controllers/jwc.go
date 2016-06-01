package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface.git/jwc_api/models"
)

// Operations about object
type JwcController struct {
	beego.Controller
}

// @Title create
// @Description create object
// @Param	body		body 	models.jwc	true		"The object content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (jwc *JwcController) Post() {
	user := jwc.GetString("username")
	pass := jwc.GetString("password")
	vrcode := jwc.GetString("vrcode")
	resView := models.Getsp()
	VS := resView["VIEWSTATE"]
	VG := resView["VIEWSTATEGENERATOR"]
	_ = models.Post(models.Get_Client(), user, pass, vrcode, VS, VG, models.GetT_cookie())
	flag, str := models.Login_OK(models.Client)
	if flag {
		jwc.Ctx.WriteString(str)
	} else {
		jwc.Ctx.WriteString("Failed")
	}
}
