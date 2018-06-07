package main

import (
	"io"
	"github.com/labstack/echo"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/middleware"
)

type LangRender struct {
}

func (l *LangRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, err := pongo2.DefaultSet.FromFile(name + ".html")
	if err != nil {
		return err
	}

	return t.ExecuteWriter(data.(pongo2.Context), w)
}
func main()  {
	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Gzip())
	app.Static("/", "static")
	app.Renderer = &LangRender{}
	app.GET("/", func(c echo.Context) error {
		return c.Render(200, "view/index", pongo2.Context{

		})
	})

	//app.GET("/login", func(c echo.Context) error {
	//	return c.Render(200, "view/login", pongo2.Context{
	//
	//	})
	//})
	//
	//app.GET("/sign-up", func(c echo.Context) error {
	//	return c.Render(200, "view/sign-up", pongo2.Context{
	//
	//	})
	//})

	app.Start(":8000")
}
