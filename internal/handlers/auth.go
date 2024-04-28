package handlers

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/dizars1776/library-lite/internal/store"
	"github.com/dizars1776/library-lite/internal/templates"
	"github.com/dizars1776/library-lite/internal/templates/components"
)

type AuthHandler struct {
	store *store.Store
}

func NewAuthHandler(store *store.Store) *AuthHandler {
	return &AuthHandler{store: store}
}

func (h *AuthHandler) Login(c echo.Context) error {
	loginComp := components.LoginForm()
	authIndex := templates.AuthIndex(loginComp, "Log In - lIlI")
	authIndex.Render(context.Background(), os.Stdout)
	return authIndex.Render(context.Background(), c.Response().Writer)
}

func (h *AuthHandler) Singup(c echo.Context) error {
	signupComp := components.SignupForm()
	authIndex := templates.AuthIndex(signupComp, "Sing Up - lIlI")
	authIndex.Render(context.Background(), os.Stdout)
	return authIndex.Render(context.Background(), c.Response().Writer)
}

// func Logout(c echo.Context) error {
//
// }
