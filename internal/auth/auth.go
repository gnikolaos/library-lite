package auth

import (
	"net/http"

	"github.com/dizars1776/library-lite/internal/services"
	"github.com/dizars1776/library-lite/internal/store"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	store *store.Store
}

func NewAuth(store *store.Store) *Auth {
	return &Auth{store: store}
}

func (a *Auth) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	// rememberMe := c.FormValue("rememberMe")
	// TODO: add JWT

	var user services.User

	err := a.store.DB().QueryRow("SELECT id, email, name, surname, password, role FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.Password, &user.Role)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.String(http.StatusUnauthorized, "- Wrong credentials, please try again")
	}

	c.Set("user", user)
	// TODO: get the user role from the user service
	c.Response().Header().Set("HX-Redirect", "/"+user.Role)

	return nil
}

// TODO: role Middleware
func (a *Auth) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return c.Redirect(http.StatusSeeOther, "/auth/login")
		}
	}
}
