package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/dizars1776/library-lite/internal/auth"
	"github.com/dizars1776/library-lite/internal/services"
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

func (h *AuthHandler) GetLogin(c echo.Context) error {
	loginComp := components.LoginForm()
	authIndex := templates.AuthIndex(loginComp, "Log In - lIlI")
	authIndex.Render(context.Background(), os.Stdout)
	return authIndex.Render(context.Background(), c.Response().Writer)
}

func (h *AuthHandler) GetSingup(c echo.Context) error {
	signupComp := components.SignupForm()
	authIndex := templates.AuthIndex(signupComp, "Sing Up - lIlI")
	authIndex.Render(context.Background(), os.Stdout)
	return authIndex.Render(context.Background(), c.Response().Writer)
}

func (h *AuthHandler) PostLogin(c echo.Context) error {
	// TODO: sanitize the values, add some checks
	email := c.FormValue("email")
	password := c.FormValue("password")
	rememberMe := c.FormValue("rememberMe") == "on"

	userService := services.NewUserService(h.store)

	loggedInUser, err := userService.LoginUser(email, password)
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusUnauthorized, "Invalid Credentials!")
	}

	err = auth.GenerateTokensAndSetCookies(loggedInUser, rememberMe, c)
	if err != nil {
		return c.HTML(http.StatusUnauthorized, "Token generation failed!")
	}

	userRole, err := userService.GetRoleName(loggedInUser.RoleId)
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusUnauthorized, "Invalid user role.")
	}

	c.Response().Header().Add("HX-Redirect", "/"+userRole)

	return nil
}

func (h *AuthHandler) GetLogout(c echo.Context) error {
	accessTokenCookie, err := c.Cookie(auth.AccessTokenCookieName)
	if err == nil {
		accessTokenCookie.Path = "/"
		accessTokenCookie.MaxAge = -1
		c.SetCookie(accessTokenCookie)
	}

	refreshTokenCookie, err := c.Cookie(auth.RefreshTokenCookieName)
	if err == nil {
		refreshTokenCookie.Path = "/"
		refreshTokenCookie.MaxAge = -1
		c.SetCookie(refreshTokenCookie)
	}

	c.Response().Header().Add("HX-Redirect", "/")

	return nil
}
