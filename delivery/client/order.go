package main

import (
	"github.com/gopherjs/jquery"
	"github.com/gopherjs/gopherjs/js"
)

func HanldeOrderBuy() {
	jquery.NewJQuery("form-order-buy").On(jquery.SUBMIT, func(e jquery.Event) {
		e.PreventDefault()
		amount := jquery.NewJQuery("#amount-buy").Val()
		price := jquery.NewJQuery("#price-buy").Val()

		posting := jquery.Post("/order", js.M{
			"symbol":    "USDEUR",
			"executedQty": amount,
			"price":     price,
			"type":      0,
		})

		posting.Done(func(data *js.Object) {
			if data.Get("Success").Bool() {

			}
		})
	})
}
