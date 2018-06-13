package main

import "github.com/gopherjs/jquery"

func ShowErr(id string) {
	jquery.NewJQuery(id).RemoveAttr("style")
}

func HideErr(id string) {
	jquery.NewJQuery(id).SetAttr("style", "display: none;")
}
