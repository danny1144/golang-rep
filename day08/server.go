package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello: " + request.GetValue()
	return nil
}
func main() {

	_ = RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listening tcp error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error :", err)
		}
		go rpc.ServeConn(conn)
	}

}
