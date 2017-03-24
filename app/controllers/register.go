package controllers

import (
	"github.com/revel/revel"
)

type Register struct {
	*revel.Controller
}

func (c Register) GetRegister() revel.Result {
	return c.Render()
}

func (c Register) DoRegister() revel.Result {
	txtfname := c.Params.Get("txtfname")
	c.Validation.Required(txtfname).Message("Please enter a username")

	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Register.GetRegister)
	}
	//models.DoRegistration()
	//fmt.Println(c.Params.Values)
	return nil
}
