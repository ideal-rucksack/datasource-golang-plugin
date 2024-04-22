package rdb

type commandFactory func(Client) Command

var commandRegistry = map[string]commandFactory{}

func RegisterCommand(name string, factory commandFactory) {
	commandRegistry[name] = factory
}

func GetCommand(action string, client Client) Command {
	if factory, ok := commandRegistry[action]; ok {
		return factory(client)
	}
	return nil
}
