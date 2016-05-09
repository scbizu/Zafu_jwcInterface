package controllers

import (
	"jwc_api/models"
	"encoding/json"
//	"net/http"
	"github.com/astaxie/beego"
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
func (o *JwcController) Post() {
//	c:=&http.Client{}

	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}


