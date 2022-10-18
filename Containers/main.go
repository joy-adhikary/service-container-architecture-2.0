package containers

type Containers struct {
	Services map[string]interface{}
}

func (c *Containers) Set(name string, class interface{}) {
	c.Services[name] = class
}

func (c *Containers) Get(name string) interface{} {
	return c.Services[name]
}

func (c *Containers) Delete(name string) {
	delete(c.Services, name)
}
