package main

import "net/rpc"

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	// 满足go语言的RPC规则的，两个可序列话参数，第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法
	Hello(request string, reply *string) error
}

// 注册服务
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
