package main

import (
	"github.com/novliang/compass/aboriginal"
	"github.com/novliang/compass/compass"
)

func main() {
	c := compass.New()
	c.Run(aboriginal.NewServer())
}
