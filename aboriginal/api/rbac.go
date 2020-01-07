package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

type RbacConfig struct {
	Guarder Guarder
}

type GuardParams map[string]string

type Guarder interface {
	CheckAccess(string, string, GuardParams) bool
}

func Rbac(r *RbacConfig) echo.MiddlewareFunc {
	//Load Db instance
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			userId := "guest"
			user := context.Get("user")
			if user != nil {
				token := user.(*jwt.Token)
				claims := token.Claims.(*JwtCustomerClaims)
				userId = strconv.Itoa(claims.Uid)
				productId := strconv.Itoa(claims.ProductId)
				if productId != "" {
					userId = fmt.Sprintf("%s_%s", userId, productId)
				}
			}
			if !r.Guarder.CheckAccess(userId, context.Path(), GuardParams{}) {
				return echo.NewHTTPError(401, "forbidden")
			}
			return next(context)
		}
	}
}
