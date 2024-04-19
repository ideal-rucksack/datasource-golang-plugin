package rdb

type Command interface {
	Execute() error
}
