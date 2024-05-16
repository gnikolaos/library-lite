package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dizars1776/library-lite/internal/services"
	"github.com/golang-jwt/jwt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

var (
	jwtSecretKey          string
	accessTokenCookieName string
)

func init() {
	accessTokenCookieName = os.Getenv("ACCESSTOKENCOOKIENAME")
	jwtSecretKey = os.Getenv("JWTSECRETKEY")
}

func GetJWTSecret() string {
	return jwtSecretKey
}

type Claims struct {
	jwt.StandardClaims
}

func GenerateTokensAndSetCookies(user *services.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)

	return nil
}

func generateAccessToken(user *services.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user *services.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "lily-app",
			Subject:   user.Name,
			Audience:  strconv.Itoa(user.RoleId),
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

// Purpose: save the user's name and other details later on
func setUserCookie(user *services.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = expiration
	cookie.Path = "/"

	c.SetCookie(cookie)
}

// Executes when user tries to access a protected path
func JWTErrorHandler(c echo.Context, err error) error {
	fmt.Println("Unauthenticated request.", err)
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

// TODO: role Middleware
// func (u *Auth) Middleware() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			return c.Redirect(http.StatusSeeOther, "/auth/login")
// 		}
// 	}
// }

// Credits: https://webdevstation.com/posts/user-authentication-with-go-using-jwt-token/
