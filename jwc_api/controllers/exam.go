package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface.git/jwc_api/jwcpkg"
	"github.com/scbizu/Zafu_jwcInterface.git/jwc_api/models"
)

type ExamController struct {
	beego.Controller
}

func (this *ExamController) Get() {
	c := models.Client
	exam := models.GetExaminfo(c)
	examInfo := jwcpkg.FetchExam(exam)
	this.Data["json"] = examInfo
	this.ServeJSON(true)
}
