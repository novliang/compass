package aboriginal

import (
	"github.com/novliang/yh_user/aboriginal/api"
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

	if config.Routers != nil {
		config.Routers(c.Engine)
	}

	return c
}

