package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Context struct {
	echo.Context
}

func (c *Context) Out(i interface{}) error {

	//New Encoder
	e := json.NewEncoder(c.Context.Response())

	//Get Header
	header := c.Response().Header()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	}

	//Set Header Code
	c.Response().WriteHeader(http.StatusOK)

	//Appointment Response
	ar := new(Response)
	ar.Code = 0
	ar.Message = "success"
	ar.Data = i

	return e.Encode(ar)
}
