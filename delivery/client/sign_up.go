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

		if password != password2 || len(password) < 6 {
			handlePasswordErr(password, password2)
			return
		}

		posting := jquery.Post("/sign-up", js.M{
			"email":    email,
			"fullname": fullname,
			"password": password,
		})

		posting.Done(func(data *js.Object) {
			if data.Get("Success").String() == "false"{
				handlePostingErr(data)
			} else {
				//jquery.NewJQuery("#modal-sign-up").ToggleClass("is-active")
			}
		})

	})
}

func CloseModalSignUp()  {
	jquery.NewJQuery("#close-sign-up").On(jquery.CLICK, func(e jquery.Event) {
		//jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery("#modal-sign-up").ToggleClass("is-active")
	})
}

func ActiveModalSignUp() {
	jquery.NewJQuery("#sign-up").On(jquery.CLICK, func(e jquery.Event) {
		HideErr("#email-is-already")
		HideErr("#not-match-pass")
		HideErr("#length-pass")

		jquery.NewJQuery("#email").SetVal("")
		jquery.NewJQuery("#password").SetVal("")
		jquery.NewJQuery("#password2").SetVal("")
		jquery.NewJQuery("#fullname").SetVal("")

		jquery.NewJQuery(e.CurrentTarget).ToggleClass("is-active")
		jquery.NewJQuery("#modal-sign-up").ToggleClass("is-active")
	})
}

func handlePasswordErr(password string, password2 string) {
	if jquery.NewJQuery("#email-is-already").Attr("style") == "" {
		HideErr("#email-is-already")
		//println("Null style")
	}
	if password != password2 {
		ShowErr("#not-match-pass")
	} else {
		HideErr("#not-match-pass")
	}

	if len(password) < 6 {
		ShowErr("#length-pass")
	} else {
		HideErr("#length-pass")
	}
}

func handlePostingErr(data *js.Object) {
	status := data.Get("Status").String()
	if status == "Email is already in use" {
		ShowErr("#email-is-already")
	}
}




