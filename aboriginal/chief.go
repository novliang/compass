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
		ip := argSlice[0]
		port := argSlice[1]
		address = ip.(string) + ":" + port.(string)
	}
	err := c.Engine.LoadConfig();
	if err != nil {
		return err
	}
	return c.Engine.Start(address)
}
