package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen tcp error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error")
	}
	rpc.ServeConn(conn)

}
