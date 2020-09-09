package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func sayHi(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(200)
	io.WriteString(wr, "hello world!")
}
func main() {

	http.HandleFunc("/", sayHi)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server listen localhost:9090")

}
