package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString(beego.AppConfig.String("custom"))
}

func (c *UserController) Index() {
	c.Ctx.WriteString("Index")
}
