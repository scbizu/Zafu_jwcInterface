package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scbizu/Zafu_jwcInterface.git/jwc_api/jwcpkg"
	"github.com/scbizu/Zafu_jwcInterface.git/jwc_api/models"
)

type CourseController struct {
	beego.Controller
}

func (this *CourseController) Get() {
	c := models.Client
	str := models.GetCourseData(c)
	data := jwcpkg.Fetchclass(str)
	this.Ctx.WriteString(data)
}
