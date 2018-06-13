package main

import (
	"io"
	"github.com/labstack/echo"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/middleware"
	"github.com/go-pg/pg"
	"projects/trading/repository"
	"projects/trading/models"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"projects/trading/usecase"
	"time"
	"fmt"
	"github.com/satori/go.uuid"
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
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("langdethuong"))))
	app.Renderer = &LangRender{}
	app.GET("/", func(context echo.Context) error {
		return context.Render(200, "delivery/view/index", pongo2.Context{

		})
	})

	app.POST("/sign-up", func(context echo.Context) error {
		print("Run")
		u := &models.User{}
		if err := context.Bind(u); err != nil {
			println(err.Error())
			return err
		}
		println(u.Password)
		var user repository.UserRepository
		user.User = u
		if user.IsAlreadyAccount(db) {
			resStr := map[string]interface{}{
				"Success": "false",
				"Status":  "Email is already in use",
			}
			return context.JSON(200, resStr)
		}

		user.SignUpAccount(db)
		resStr := map[string]interface{}{
			"Success": "true",
			"Status":  "OK",
			"Name":    u.Fullname,
		}
		return context.JSON(200, resStr)
	})

	app.POST("/sign-in", func(context echo.Context) error {
		u := &models.User{}
		if err := context.Bind(u); err != nil {
			println(err.Error())
			return err
		}

		var user repository.UserRepository
		user.User = u

		if !user.IsAlreadyAccount(db) {
			res := map[string]interface{}{
				"Success": false,
			}
			return context.JSON(200, res)
		}
		usecase.SetSession(context, u)
		res := map[string]interface{}{
			"Success": true,
		}
		return context.JSON(200, res)
	})

	app.POST("/order", func(context echo.Context) error {
		sess, _ := session.Get("session", context)
		uuidStr := fmt.Sprint(sess.Values["ID"])
		id, _ := uuid.FromString(uuidStr)

		order := &models.Order{}
		if err := context.Bind(order); err != nil {
			println(err.Error())
			return err
		}

		balance := repository.SelectBalance(db, id)
		if order.Type == 0 {
			order.OriginQty = balance.USD
		} else {
			order.OriginQty = balance.EUR
		}

		if order.ExecutedQty > order.OriginQty {
			res := map[string]interface{}{
				"Success": false,
			}
			return context.JSON(200, res)
		}
		order.Time = time.Now()
		order.IsWorking = true
		order.UserID = id
		return context.JSON(200, id)
	})

	app.Start(":8000")
}
