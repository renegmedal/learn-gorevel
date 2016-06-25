package controllers

import (
	"github.com/revel/revel"
)

type Orders struct {
	*revel.Controller /// to get access to revel's controller methods, etc..
}

// localhost:9000/orders/create

// c Orders matches /views/orders
// Create() matches views/orders/create.html
func (c Orders) Create() revel.Result { // returns so caller recognizes that it is a revel result
	return c.Render()
}
