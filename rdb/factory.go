package rdb

type commandFactory func(Client) Command

var commandRegistry = map[string]commandFactory{}

func registerCommand(name string, factory commandFactory) {
	commandRegistry[name] = factory
}

func GewCommand(api string, client Client) Command {
	if factory, ok := commandRegistry[api]; ok {
		return factory(client)
	}
	return nil
}
