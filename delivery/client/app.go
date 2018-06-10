package main

import "github.com/gopherjs/jquery"

func main() {

	ActiveNavbarMenu()
	ActiveModalSignIn()
	ActiveModalSignUp()
	CloseModalSignIn()
	CloseModalSignUp()
	HandleSignUp()
}

func ActiveNavbarMenu()  {
	jquery.NewJQuery(".navbar-burger").On(jquery.CLICK, func(e jquery.Event) {
		jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery(".navbar-menu").ToggleClass("is-active")
	})
}



func ActiveModalSignIn()  {
	jquery.NewJQuery("#sign-in").On(jquery.CLICK, func(e jquery.Event) {

		jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery("#modal-sign-in").ToggleClass("is-active")
	})
}

func CloseModalSignUp()  {
	jquery.NewJQuery("#close-sign-up").On(jquery.CLICK, func(e jquery.Event) {
		//jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery("#modal-sign-up").ToggleClass("is-active")
	})
}

func CloseModalSignIn()  {
	jquery.NewJQuery("#close-sign-in").On(jquery.CLICK, func(e jquery.Event) {
		//jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery("#modal-sign-in").ToggleClass("is-active")
	})
}


