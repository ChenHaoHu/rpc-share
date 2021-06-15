package main

import (
	"flag"
	"rpcx/server/api"

	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "127.0.0.1:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	// s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(api.Arith), "")
	s.Serve("tcp", *addr)
}
