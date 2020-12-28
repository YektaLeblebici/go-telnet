package telnet

import "net"

type Context interface {
	Logger() Logger

	InjectLogger(Logger) Context
	InjectConn(net.Conn) Context
}

type internalContext struct {
	logger Logger
	conn   net.Conn
}

func NewContext() Context {
	ctx := internalContext{}

	return &ctx
}

func (ctx *internalContext) Logger() Logger {
	return ctx.logger
}

func (ctx *internalContext) InjectLogger(logger Logger) Context {
	ctx.logger = logger

	return ctx
}

func (ctx *internalContext) InjectConn(c net.Conn) Context {
	ctx.conn = c

	return ctx
}
