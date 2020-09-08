package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error")
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)

	if err != nil {
		log.Fatal("call method error:", err)
	}
	fmt.Println(reply)

}
