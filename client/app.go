package main

import "github.com/gopherjs/jquery"

func main() {
	jquery.NewJQuery(".navbar-burger").On(jquery.CLICK, func(e jquery.Event) {
		jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery(".navbar-menu").ToggleClass("is-active")
	})

	jquery.NewJQuery("#login").On(jquery.CLICK, func(e jquery.Event) {
		jquery.NewJQuery("modal").Show()
	})

}
