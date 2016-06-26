package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/revel/revel"
	"github.com/vansimke/restaurant-site/app/models"
)

// to have 64-bit randome generator
var rndGen = rand.New(rand.NewSource(time.Now().Unix()))

const chance = 0.5

type Orders struct {
	*revel.Controller
	winner bool
}

// Order pointer is used to store state through out the request call
// But each new request resets the state of the value of winner
// but it less performant
func (c *Orders) randomDrawing() revel.Result {
	if rndGen.Float32() < chance {
		c.winner = true
	}
	return nil
}

func (c Orders) Create() revel.Result {
	return c.Render()
}

func (c Orders) GetPayment(orderId int) revel.Result {
	println("The order ID: ", orderId)
	return c.RenderTemplate("orders/payment.html")
}

func (c Orders) ApiCreate() revel.Result {
	var order models.Order

	if c.winner {
		println("You have won a discount!!!!!")
	}

	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&order)
	fmt.Printf("The order data: %v\n", order)
	return c.RenderText("OK")
}

func init() {
	// intercept method randowDrawing before it is called
	revel.InterceptMethod((*Orders).randomDrawing, revel.BEFORE)
}
