package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "spletnakoda.si"
	c.Data["Email"] = "gorazd@spletnakoda.si"
	c.TplName = "index.tpl"
}
