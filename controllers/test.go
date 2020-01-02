package controllers

import (
	"github.com/GorazdVeselic/beego-swager/models"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

// @Title TestEndpoint
// @Description Tests the API
// @Success 200 {object} models.Test
// @Failure 403 body is empty
// @router / [get]
func (t *TestController) Test() {
	Response := models.TestFunction()
	t.Data["json"] = Response
	t.ServeJSON()
}
