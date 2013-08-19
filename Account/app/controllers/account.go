package controllers

import (
	"github.com/robfig/revel"
	//"fmt"
	"github.com/jsli/revel-in-action/Account/app/models"
)

type Account struct {
	*revel.Controller
}

func (c Account) Index() revel.Result {
	return c.Render()
}

func (c Account) GetLogin() revel.Result {
	return c.Render()
}

func (c Account) PostLogin() revel.Result {
	return c.Render()
}

func (c Account) Logout() revel.Result {
	return c.Render()
}

func (c Account) GetRegister() revel.Result {
	return c.Render()
}

/*
 * regUser is a struct's name in the template
 * see {{with $field := field "regUser.Field" .}} in template
 */
func (c Account) PostRegister(regUser *models.RegUser) revel.Result {
	//step 0: check user is exist or not
	regUser.Validate(c.Validation)

	//step 1: validation
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Account.GetRegister)
	}

	//step 2: save user
	regUser.SaveUser(regUser)

	//step 3: save cookie, flash or session
	c.Session["user"] = regUser.UserName
	c.Flash.Success("Welcome, " + regUser.UserName)

	//step 4: rediret
	return c.Redirect(Account.Index)
}

