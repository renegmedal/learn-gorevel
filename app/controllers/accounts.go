package controllers

import (
	"fmt"
	"pizzastore/app/models"

	"github.com/revel/revel"
)

type Accounts struct {
	*revel.Controller
}

func (a Accounts) Create() revel.Result {
	return a.Render()
}

func (a Accounts) CreatePost() revel.Result {
	var account models.Account
	a.Params.Bind(&account, "account")

	fmt.Printf("Account info: %v\n", account)

	return a.RenderTemplate("accounts/create.html")
}
