package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dizars1776/library-lite/internal/services"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

var (
	AccessTokenCookieName  string
	RefreshTokenCookieName string
	jwtSecretKey           string
	jwtRefreshSecretKey    string
)

func init() {
	AccessTokenCookieName = os.Getenv("ACCESSTOKENCOOKIENAME")
	RefreshTokenCookieName = os.Getenv("REFRESHTOKENCOOKIENAME")
	jwtSecretKey = os.Getenv("JWTSECRETKEY")
	jwtRefreshSecretKey = os.Getenv("JWTREFRESHSECRETKEY")
}

type jwtCustomClaims struct {
	Name       string `json:"name"`
	RememberMe bool   `json:"rememberMe"`
	jwt.RegisteredClaims
}

func GetJWTSecret() string {
	return jwtSecretKey
}

func GetJWTRefreshSecret() string {
	return jwtRefreshSecretKey
}

func GetJWTCustomClaims(c echo.Context) jwt.Claims {
	return new(jwtCustomClaims)
}

func GenerateTokensAndSetCookies(user *services.User, rememberMe bool, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user, rememberMe)
	if err != nil {
		return err
	}

	setTokenCookie(AccessTokenCookieName, accessToken, exp, c)
	refreshToken, exp, err := generateRefreshToken(user, rememberMe)
	if err != nil {
		return err
	}
	setTokenCookie(RefreshTokenCookieName, refreshToken, exp, c)

	return nil
}

func generateAccessToken(user *services.User, rememberMe bool) (string, time.Time, error) {
	accessTokenExp := 60 * time.Minute

	// If rememberMe = true keep the user logged in for 1 month
	if rememberMe {
		accessTokenExp = 24 * 30 * time.Hour
	}

	expirationTime := time.Now().Add(accessTokenExp)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()), rememberMe)
}

func generateRefreshToken(user *services.User, rememberMe bool) (string, time.Time, error) {
	refreshTokenExp := 24 * time.Hour

	// If rememberMe = true keep the user logged in for 1 month
	if rememberMe {
		refreshTokenExp = 24 * 30 * time.Hour
	}

	expirationTime := time.Now().Add(refreshTokenExp)

	return generateToken(user, expirationTime, []byte(GetJWTRefreshSecret()), rememberMe)
}

func generateToken(user *services.User, expirationTime time.Time, secret []byte, rememberMe bool) (string, time.Time, error) {
	claims := &jwtCustomClaims{
		Name:       user.Name,
		RememberMe: rememberMe,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "lily-app",
			Subject:   user.Email,
			Audience:  []string{strconv.Itoa(user.RoleId)},
			ExpiresAt: jwt.NewNumericDate(expirationTime),
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

// Executes when user tries to access a protected path
func JWTErrorHandler(c echo.Context, err error) error {
	fmt.Println("Unauthenticated request.", err)
	return c.Redirect(http.StatusSeeOther, c.Echo().Reverse("userLogInForm"))
}

func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken, ok := c.Get("user").(*jwt.Token)
		if !ok {
			fmt.Println("No user token in context")
			return next(c)
		}

		claims, ok := userToken.Claims.(*jwtCustomClaims)
		if !ok || !userToken.Valid {
			fmt.Println("Invalid user token")
			fmt.Printf("Claims: %+v\n", userToken.Claims)
			return next(c)
		}

		if time.Until(claims.ExpiresAt.Time) < 15*time.Minute {
			rc, err := c.Cookie(RefreshTokenCookieName)
			if err == nil && rc != nil {
				refreshClaims := &jwtCustomClaims{}
				tkn, err := jwt.ParseWithClaims(rc.Value, refreshClaims, func(token *jwt.Token) (interface{}, error) {
					return []byte(GetJWTRefreshSecret()), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					_ = GenerateTokensAndSetCookies(&services.User{}, refreshClaims.RememberMe, c)
				}
			}
		}

		return next(c)
	}
}

// Credits: https://webdevstation.com/posts/user-authentication-with-go-using-jwt-token/
