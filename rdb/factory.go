package rdb

type CommandFactory[T any] func(Client) Command[T]

type CommandFactoryRegistry[T any] struct {
	factories map[string]CommandFactory[T]
}

func NewCommandFactoryRegistry[T any]() *CommandFactoryRegistry[T] {
	return &CommandFactoryRegistry[T]{
		factories: make(map[string]CommandFactory[T]),
	}
}

func (r *CommandFactoryRegistry[T]) Register(name string, factory CommandFactory[T]) {
	r.factories[name] = factory
}

func (r *CommandFactoryRegistry[T]) GetCommand(name string, client Client) (Command[T], bool) {
	factory, ok := r.factories[name]
	if !ok {
		return nil, false
	}
	c := factory(client)
	return c, true
}
