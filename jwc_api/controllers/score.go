package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface/jwc_api/jwcpkg"
	"github.com/scbizu/Zafu_jwcInterface/jwc_api/models"
)

type ScoreController struct {
	beego.Controller
}

func (this *ScoreController) Get() {
	info := models.GetScoreinfo(models.Client)
	vs := models.GetscoreVs(info)
	vg := models.GetscoreVg(info)
	beego.Debug(models.GetT_cookie())
	data := models.FindOutScore(models.Client, vs, vg, "", "", "", models.GetT_cookie())
	beego.Debug(strings.TrimSpace(data))
	// tbody := jwcpkg.DirectTbody(strings.TrimSpace(data))
	Scoreinfo := jwcpkg.FetchScoreTD(data)
	this.Data["json"] = Scoreinfo
	this.ServeJSON(true)
}
