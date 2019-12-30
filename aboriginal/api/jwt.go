package api

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/novliang/compass/app/service"
	"github.com/spf13/viper"
	"strings"
)

type JwtCustomerClaims struct {
	Uid   int    `json:"uid"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Jwt() echo.MiddlewareFunc {
	jwtConfig := viper.Get("jwt").(map[string]interface{})
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			skippers := jwtConfig["skipper"].([]interface{})
			skippers = append(skippers, "GET|/ping")
			for _, v := range skippers {
				vs := v.(string)
				sv := strings.Split(vs, "|")
				if len(sv) < 2 {
					panic("some thing wrong with jwt skipper config")
				}
				if sv[0] == c.Request().Method {
					if sv[1] == c.Path() {
						return true
					}
				}
			}

			if mapset.NewSetFromSlice(skippers).Contains(c.Path()) {
				return true
			}
			return false
		},
		SigningKey: []byte(jwtConfig["secret"].(string)),
		Claims:     &JwtCustomerClaims{},
	})
}

func GenerateToken(claims *JwtCustomerClaims) (t string, err *service.ServiceError) {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.

	jwtConfig := ApiConfigurator.GetStringMap("jwt")

	t, tokenErr := token.SignedString([]byte(jwtConfig["secret"].(string)))
	if tokenErr != nil {
		return "", service.NewServiceError(tokenErr.Error())
	}
	return
}
