package controllers

import (
	"github.com/revel/revel"
)

type Accounts struct {
	*revel.Controller
}

func (a Accounts) Create() revel.Result {
	return a.Render()
}
