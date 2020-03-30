package aboriginal

type Engine interface {
	Start(string) error
	LoadConfig() error
}

type Chief struct {
	Engine Engine
}

func (c *Chief) Run(args ...interface{}) error {
	address := ""
	if len(args) > 0 {
		argSlice := args[0].([]interface{})
		if len(argSlice) == 2 {
			ip := argSlice[0]
			port := argSlice[1]
			address = ip.(string) + ":" + port.(string)
		}
	}
	return c.Engine.Start(address)
}
