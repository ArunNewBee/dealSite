package controllers

import (
	"fmt"

	"github.com/revel/revel"
)

type Register struct {
	*revel.Controller
}

func (c Register) GetRegister() revel.Result {
	return c.Render()
}

func (c Register) DoRegister() revel.Result {
	fmt.Println("hello")
	fmt.Println(c.Params.Values)
	return nil
}
