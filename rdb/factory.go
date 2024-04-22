package rdb

type CommandFactory func(Client) Command

type CommandFactoryRegistry struct {
	factories map[string]CommandFactory
}

func NewCommandFactoryRegistry() *CommandFactoryRegistry {
	return &CommandFactoryRegistry{
		factories: make(map[string]CommandFactory),
	}
}

func (r *CommandFactoryRegistry) Register(name string, factory CommandFactory) {
	r.factories[name] = factory
}

func (r *CommandFactoryRegistry) GetCommand(name string, client Client) (Command, bool) {
	factory, ok := r.factories[name]
	if !ok {
		return nil, false
	}
	c := factory(client)
	return c, true
}
