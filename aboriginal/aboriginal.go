package aboriginal

import (
	"github.com/novliang/compass/aboriginal/api"
)

type RoutersInjection func(engine Engine)

type AboriginalConfig struct {
	Engine  Engine
	Routers RoutersInjection
}

func DefaultEngine() interface{} {
	return api.Engine()
}

func NewServer() *Chief {
	c := &Chief{
		Engine: DefaultEngine().(Engine),
	}
	return c
}

func NewServerByConfig(config AboriginalConfig) *Chief {

	c := &Chief{}

	if config.Engine != nil {
		c.Engine = config.Engine
	} else {
		c.Engine = DefaultEngine().(Engine)
	}

	err := c.Engine.LoadConfig()
	if err != nil {
		panic(err)
	}

	if config.Routers != nil {
		config.Routers(c.Engine)
	}

	return c
}
