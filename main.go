package main

import (
	"github.com/novliang/yh_user/aboriginal"
	"github.com/novliang/yh_user/app"
	"github.com/novliang/yh_user/compass"
)

func main() {
	c := compass.New()
	c.Run(aboriginal.NewServerByConfig(aboriginal.AboriginalConfig{
		Routers: app.RouterInjection(),
	}), "192.168.1.224", "8098")
}
