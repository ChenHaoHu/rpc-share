package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	api2 "rpcx/base/server/api"
)

var addr = flag.String("addr", "127.0.0.1:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	// s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(api2.Arith), "")
	s.Serve("tcp", *addr)
}
