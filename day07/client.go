package main

import (
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DiaHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

//客户端调用RPC方法
func main() {
	client, err := DiaHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error")
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal("call error: ", err)
	}

	println("reply: ", reply)
}
