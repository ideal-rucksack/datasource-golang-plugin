package rdb

type Command interface {
	// Execute 执行命令
	Execute() (any, error)
	// Notify 回调通知调用方
	// args 为执行命令的参数
	// payload 为执行结果作为POST请求的body(请求体)
	Notify(args ExecCommand, payload any) error
}

type NotifyRequest struct {
	status  int
	error   *string
	payload any
}
