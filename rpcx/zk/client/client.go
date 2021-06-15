package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/smallnest/rpcx/client"
	example "zk/server/api"
)

var (
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	d, _ := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {

		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}

}
