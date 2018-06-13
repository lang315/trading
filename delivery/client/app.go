package main

import "github.com/gopherjs/jquery"

func main() {
	ActiveNavbarMenu()
	SignUp()
	SignIn()
	HanldeOrderBuy()
}

func ActiveNavbarMenu()  {
	jquery.NewJQuery(".navbar-burger").On(jquery.CLICK, func(e jquery.Event) {
		jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery(".navbar-menu").ToggleClass("is-active")
	})
}

func SignUp()  {
	ActiveModalSignUp()
	CloseModalSignUp()
	HandleSignUp()
}

func SignIn()  {
	ActiveModalSignIn()
	CloseModalSignIn()
	HandleSignIn()
}



