package main

import (
	"flag"
	"log"

	"github.com/aintsashqa/go-project-template/delivery/jsonrpc"
	"github.com/aintsashqa/go-project-template/handler"
)

var (
	dev = flag.Bool("dev", false, "development mode")
)

func init() {
	flag.Parse()
}

func main() {
	srv := jsonrpc.New(jsonrpc.Config{
		Addr:     "127.0.0.1:3000",
		Endpoint: "/api",
	})

	if err := srv.
		RegisterName("Todo", handler.NewTodoHandler()).
		ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
