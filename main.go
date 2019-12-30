package main

import (
	"github.com/novliang/compass/aboriginal"
	"github.com/novliang/compass/app"
	"github.com/novliang/compass/compass"
)

func main() {
	c := compass.New()
	c.Run(aboriginal.NewServerByConfig(aboriginal.AboriginalConfig{
		Routers: app.RouterInjection(),
	}), "192.168.1.224", "8098")
}
