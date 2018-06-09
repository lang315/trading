package main

import (
	"github.com/gopherjs/jquery"
	"github.com/gopherjs/gopherjs/js"
)

func HandleSignUp() {
	jquery.NewJQuery("#form-sign-up").On(jquery.SUBMIT, func(e jquery.Event) {
		e.PreventDefault()
		password := jquery.NewJQuery("#password").Val()
		password2 := jquery.NewJQuery("#password2").Val()
		email := jquery.NewJQuery("#email").Val()
		fullname := jquery.NewJQuery("#fullname").Val()

		posting := jquery.Post("/sign-up", js.M{
			"email":    email,
			"fullname": fullname,
			"password": password,
		})

		posting.Done(func(data *js.Object) {
			println(data.Get("Email").String())
		})

		if password != password2 {
			jquery.NewJQuery("#not-match-pass").RemoveAttr("style")

		} else
		{
			jquery.NewJQuery("#not-match-pass").SetAttr("style", "display: none;")
		}
		if len(password) < 6 {
			jquery.NewJQuery("#length-pass").RemoveAttr("style")
		} else {
			jquery.NewJQuery("#length-pass").SetAttr("style", "display: none;")

		}
	})
}

func ActiveModalSignUp() {
	jquery.NewJQuery("#sign-up").On(jquery.CLICK, func(e jquery.Event) {
		jquery.NewJQuery("#length-pass").SetAttr("style", "display: none;")
		jquery.NewJQuery("#not-match-pass").SetAttr("style", "display: none;")

		jquery.NewJQuery("#email").SetVal("")
		jquery.NewJQuery("#password").SetVal("")
		jquery.NewJQuery("#password2").SetVal("")
		jquery.NewJQuery("#fullname").SetVal("")

		jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery("#modal-sign-up").ToggleClass("is-active")
	})
}
