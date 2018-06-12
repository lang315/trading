package main

import (
	"io"
	"github.com/labstack/echo"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/middleware"
	"projects/trading/models"
	"projects/trading/repository"
	"github.com/go-pg/pg"
	"projects/trading/usecase"
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
func main() {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "Conghuy.315",
		Database: "trading",
	})

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Gzip())
	app.Static("/", "delivery")
	app.Renderer = &LangRender{}
	app.GET("/", func(context echo.Context) error {
		return context.Render(200, "delivery/view/index", pongo2.Context{

		})
	})

	app.POST("/sign-up", func(context echo.Context) error {
		u := &models.User{}
		if err := context.Bind(u); err != nil {
			println(err.Error())
			return err
		}

		u.Password, _ = usecase.HashPassword(u.Password)
		println(u.Password)
		var user repository.UserRepository
		user.User = u

		if user.IsAlreadyAccount(db) {
			resStr := map[string]string{
				"Success": "false",
				"Status":  "Email is already in use",
			}
			return context.JSON(200, resStr)
		}

		user.SignUpAccount(db)
		resStr := map[string]string{
			"Success": "true",
			"Status":  "OK",
			"Name":    u.Fullname,
		}
		return context.JSON(200, resStr)
	})

	app.Start(":8000")
}
