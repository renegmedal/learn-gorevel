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

// route  GET /orders/:orderId/pay
// localhost:9000/orders/12345/pay
func (o Orders) GetPayment(orderId int) revel.Result {

	println("++++ orders.GetPayment() order ID: ", orderId)

	return o.RenderTemplate("orders/payment.html") // path relative to /views/folder
}

// route GET /:controller/:orderid -> /orders/556677
func (o Orders) Info(orderId int) revel.Result {
	println("+++++ orders.OrderInfo ORDER ID: ", orderId)

	return o.RenderTemplate("App/index.html")

}

// route GET /:controller/:orderid/:action -> /orders/556677/orderinfo
func (o Orders) ProcessPayment(orderId int) revel.Result {

	println("+++++ orders.ProcessPayment ORDER ID: ", orderId)

	return o.RenderTemplate("App/index.html")
}
