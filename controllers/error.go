package controllers
import (
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}
func (c *ErrorController) Error401() {
	c.Data["content"] = "401"
	c.TplName="error.tpl"
}
func (c *ErrorController) Error403() {
	c.Data["content"] = "403"
	c.TplName="error.tpl"
}
func (c *ErrorController) Error404() {
	c.Data["content"] = "404"
	c.TplName="error.tpl"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "500"
	c.TplName = "error.tpl"
}
func (c *ErrorController) Error503() {
	c.Data["content"] = "503"
	c.TplName = "error.tpl"
}
