package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface/jwc_api/models"
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
	models.StuNo = user
	pass := jwc.GetString("password")
	vrcode := jwc.GetString("vrcode")
	//	beego.Debug(pass)
	resView := models.Getsp()
	VS := resView["VIEWSTATE"]
	VG := resView["VIEWSTATEGENERATOR"]
	cookie, c := models.Post(models.Client, user, pass, vrcode, VS, VG, models.GetT_cookie())
	flag, _ := models.Login_OK(c)
	if flag {
		models.Client = c
		models.Login_co = cookie
		beego.Debug(models.Login_co)
		jwc.Ctx.WriteString("true")
	} else {
		jwc.Ctx.WriteString("Failed")
	}
}
