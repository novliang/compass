package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/novliang/yh_user/utils"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

var ApiConfigurator *viper.Viper

type Api struct {
	*echo.Echo
}

func (api *Api) LoadConfig() error {

	w, err := os.Getwd()
	if err != nil {
		return errors.New("Can't get path wd err")
	}
	a, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return errors.New("Can't get path wd err")
	}

	var configFile = "api.toml"
	appConfigPath := filepath.Join(w, "config", configFile)
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(a, "config", configFile)
		if !utils.FileExists(appConfigPath) {
			return errors.New("Can't get db config file err")
		}
	}

	ApiConfigurator = viper.New()

	ApiConfigurator.SetConfigName("api")
	ApiConfigurator.AddConfigPath(strings.TrimRight(appConfigPath, configFile))
	err = ApiConfigurator.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("Can't get db config file err")
		} else {
			return err
		}
	}
	return nil;
}

func Engine() interface{} {

	e := echo.New()

	a := &Api{e}

	a.HTTPErrorHandler = HttpErrorHandler

	a.Use(middleware.Logger())

	a.Use(middleware.Recover())

	a.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	a.Validator = &Validator

	//Extending
	a.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			think := &Context{c}
			return handlerFunc(think)
		}
	})

	//Use gzip with level 5
	a.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	//For HA Health Check
	a.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	return a
}

func HttpErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  string
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if h, o := he.Message.(string); o {
			msg = h
		} else {
			msg = "服务器出错!"
		}
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			r := new(Response)
			r.Message = msg
			r.Code = code
			err := c.JSON(200, r)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
