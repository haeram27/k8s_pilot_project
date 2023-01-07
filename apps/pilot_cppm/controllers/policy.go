package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type PolicyController struct {
	beego.Controller
}

func (c *PolicyController) Get() {
	c.TplName = "policy.tpl"
}
