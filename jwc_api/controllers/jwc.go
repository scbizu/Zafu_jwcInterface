package controllers

import "github.com/astaxie/beego"

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
func (o *JwcController) Post() {
	//	c:=&http.Client{}

}

// func (o *JwcController) Get() {
// 	c := &http.Client{}
// 	stuid := o.Ctx.Input.Param(":stuid")
// 	password := o.Ctx.Input.Param(":password")
// 	ok := models.Login_OK(c,stuid,password)
// }
