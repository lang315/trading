package main

import (
	"github.com/gopherjs/jquery"
	"github.com/gopherjs/gopherjs/js"
)

func ActiveModalSignIn() {
	jquery.NewJQuery("#sign-in").On(jquery.CLICK, func(e jquery.Event) {
		HideErr("#email-password-incorrect")
		jquery.NewJQuery("#email-sign-in").SetVal("")
		jquery.NewJQuery("#password-sign-in").SetVal("")

		jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery("#modal-sign-in").ToggleClass("is-active")
	})
}

func CloseModalSignIn() {
	jquery.NewJQuery("#close-sign-in").On(jquery.CLICK, func(e jquery.Event) {
		jquery.NewJQuery("#modal-sign-in").ToggleClass("is-active")
	})
}

func HandleSignIn() {
	jquery.NewJQuery("#form-sign-in").On(jquery.SUBMIT, func(e jquery.Event) {
		e.PreventDefault()
		email:=jquery.NewJQuery("#email").Val()

		password := jquery.NewJQuery("#password").Val()
		println(email)
		println(password)

		posting := jquery.Post("/sign-in", js.M{
			"email":    email,
			"password": password,
		})

		posting.Done(func(data *js.Object) {
			if data.Get("Success").Bool() {
				jquery.NewJQuery("#modal-sign-in").ToggleClass("is-active")
			}else {
				println("Err")
				ShowErr("#email-password-incorrect")
			}
		})
	})
}

