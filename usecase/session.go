package usecase

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"projects/trading/models"
)

func SetSession(c echo.Context, user *models.User) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["ID"] = user.ID
	sess.Save(c.Request(), c.Response())
}
