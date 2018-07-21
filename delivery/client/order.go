package main

import (
	"github.com/gopherjs/jquery"
	"github.com/gopherjs/gopherjs/js"
)

func HandleOrderBuy() {
	jquery.NewJQuery("#form-order-buy").On(jquery.SUBMIT, func(e jquery.Event) {
		e.PreventDefault()
		amount := jquery.NewJQuery("#amount-buy").Val()
		price := jquery.NewJQuery("#price-buy").Val()

		posting := jquery.Post("/order", js.M{
			"symbol":      "USDEUR",
			"executedQty": amount,
			"price":       price,
			"type":        0,
		})

		posting.Done(func(data *js.Object) {
			if data.Get("Success").Bool() {
				if data.Get("Success").Bool() {
					obj := data.Get("Order")
					jquery.NewJQuery("#table-order-sell").On(jquery.AJAXCOMPLETE, func(e jquery.Event) {
						jquery.NewJQuery("price-order-sell").SetText(obj.Get("Price"))
					})
				} else {

				}
			}
		})
	})
}

func HanldeOrderSell() {
	jquery.NewJQuery("#form-order-sell").On(jquery.SUBMIT, func(e jquery.Event) {
		e.PreventDefault()
		amount := jquery.NewJQuery("#amount-sell").Val()
		price := jquery.NewJQuery("#price-sell").Val()

		posting := jquery.Post("/order", js.M{
			"symbol":      "USDEUR",
			"executedQty": amount,
			"price":       price,
			"type":        1,
		})

		posting.Done(func(data *js.Object) {
			if data.Get("Success").Bool() {
				obj := data.Get("Order")
				jquery.NewJQuery("#table-order-sell").On(jquery.AJAXCOMPLETE, func(e jquery.Event) {
					jquery.NewJQuery("price-order-sell").SetText(obj.Get("Price"))
				})
			} else {

			}
		})
	})
}
