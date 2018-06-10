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
			if data.Get("Success").String() == "false" {
				handlePostingErr(data)
			}
		})

	})
}

func ActiveModalSignUp() {
	jquery.NewJQuery("#sign-up").On(jquery.CLICK, func(e jquery.Event) {
		hideErrSignUp("#email-is-already")
		hideErrSignUp("#not-match-pass")
		hideErrSignUp("#length-pass")

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
		hideErrSignUp("#email-is-already")
		//println("Null style")
	}
	if password != password2 {
		showErrSignUp("#not-match-pass")
	} else {
		hideErrSignUp("#not-match-pass")
	}

	if len(password) < 6 {
		showErrSignUp("#length-pass")
	} else {
		hideErrSignUp("#length-pass")
	}
}

func handlePostingErr(data *js.Object) {
	status := data.Get("Status").String()
	if status == "Email is already in use" {
		showErrSignUp("#email-is-already")
	}
}

func showErrSignUp(id string) {
	jquery.NewJQuery(id).RemoveAttr("style")
}

func hideErrSignUp(id string) {
	jquery.NewJQuery(id).SetAttr("style", "display: none;")
}
