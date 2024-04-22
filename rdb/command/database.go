package command

import (
	"bytes"
	"encoding/json"
	"github.com/ideal-rucksack/datasource-golang-plugin/rdb"
	"net/http"
)

type DatabaseCommand struct {
	Client rdb.Client
}

func (d DatabaseCommand) Execute() (rdb.Result, error) {
	var err error
	databases, err := d.Client.Databases()
	if err != nil {
		return nil, err
	}
	result := rdb.StringsResult{Data: &databases}
	return result, nil
}

func (d DatabaseCommand) Notify(args rdb.ExecCommand, payload rdb.NotifyRequest) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", args.Webhook, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	client := &http.Client{}
	_, err = client.Do(request)
	return err
}
