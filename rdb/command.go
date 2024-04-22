package rdb

type Command[T any] interface {
	// Execute 执行命令
	Execute() (T, error)
	// Notify 回调通知调用方
	// args 为执行命令的参数
	// payload 为执行结果作为POST请求的body(请求体)
	Notify(args ExecCommand, payload NotifyRequest[T]) error
}

type NotifyRequest[T any] struct {
	Status  int     `json:"status"`
	Error   *string `json:"error"`
	Payload T       `json:"payload"`
}
