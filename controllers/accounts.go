package controllers

import (
	"bapi/models"
	"github.com/astaxie/beego"
)

// oprations for Accounts
type AccountsController struct {
	beego.Controller
}

func (c *AccountsController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// @Title Get
// @Description get Accounts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Accounts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AccountsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	v := models.GetAccountsById(idStr)
	c.Data["json"] = v
	c.ServeJSON()
}

