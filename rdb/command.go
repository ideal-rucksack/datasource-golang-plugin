package rdb

type Command interface {
	// Execute 执行命令
	Execute() (*Result, error)
	// Notify 回调通知调用方
	// args 为执行命令的参数
	// payload 为执行结果作为POST请求的body(请求体)
	Notify(args ExecCommand, payload NotifyRequest) error
}

type Result interface {
	Value() any
}

type NotifyRequest struct {
	Status  int     `json:"status"`
	Error   *string `json:"error"`
	Payload Result  `json:"payload"`
}

type StringsResult struct {
	Data *[]string
}

func (s StringsResult) Value() any {
	return s.Data
}
