package command

import "github.com/ideal-rucksack/datasource-golang-plugin/rdb"

type DatabaseCommand struct {
	client rdb.Client
}

func (d *DatabaseCommand) Execute() ([]string, error) {
	return d.client.Databases()
}
