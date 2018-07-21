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
		passHashing, _ := usecase.HashPassword(u.Password)
		println(passHashing)
		u.Password = passHashing
		var user repository.UserRepository
		user.User = u
		check, _ := user.IsAlreadyAccount(db)
		if check {
			resStr := map[string]interface{}{
				"Success": "false",
				"Status":  "Email is already in use",
				"Fullname": u.Fullname,
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
		checkEmail, userLogin := user.IsAlreadyAccount(db)
		hash,_:= usecase.HashPassword(userLogin.Password)
		checkPass:= usecase.CheckPasswordHash(userLogin.Password, hash)

		if !checkEmail || !checkPass {
			res := map[string]interface{}{
				"Success": false,
			}
			return context.JSON(200, res)
		}

		usecase.SetSession(context, u)

		res := map[string]interface{}{
			"Success": true,
			"Fullname": userLogin.Fullname,
		}
		return context.JSON(200, res)
	})

	app.POST("/order", func(context echo.Context) error {
		sess, _ := session.Get("session", context)

		if sess.Values["ID"] == nil {
			println("Err")
			res := map[string]interface{}{
				"Success": false,
			}
			return context.JSON(200, res)
		}

		uuidStr := fmt.Sprintln(sess.Values["ID"])
		id, _ := uuid.FromString(uuidStr)

		order := &models.Order{}
		if err := context.Bind(order); err != nil {
			println(err.Error())
			return err
		}

		balance := repository.SelectBalance(db, id)
		if order.Type == 0 {
			order.OriginQuantity = balance.USD
		} else {
			order.OriginQuantity = balance.EUR
		}

		if order.ExecutedQuantity > order.ExecutedQuantity {
			res := map[string]interface{}{
				"Success": false,
			}
			return context.JSON(200, res)
		}
		order.UserID = id
		order.Time = time.Now()
		order.IsWorking = true
		repository.InsertOrder(db, order)
		repository.UpdateBalance(db, balance)
		res := map[string]interface{}{
			"Success": true,
			"Order": order,
			"Balance": balance,
		}
		return context.JSON(200, res)
	})

	app.Start(":8000")
}
