package command

import (
	"fmt"
	"github.com/ideal-rucksack/datasource-golang-plugin/rdb"
)

type DatabaseCommand struct {
	Client rdb.Client
}

func (d DatabaseCommand) Execute() error {
	databases, err := d.Client.Databases()
	fmt.Printf("database: %s", databases)
	return err
}
