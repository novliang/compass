package compass

import (
	"github.com/labstack/echo/v4"
	myLog "github.com/labstack/gommon/log"
	"github.com/novliang/compass/granary"
)

type Compass struct {
	Logger echo.Logger
}

func New() (c *Compass) {
	c = &Compass{Logger: myLog.New("compass")}
	err := granary.Load()
	if err != nil {
		c.Logger.Fatal(err)
	}
	return
}

func (c *Compass) Run(s Server, args ...interface{}) {
	c.Logger.Fatal(s.Run(args));
}
