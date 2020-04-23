package main

import (
	"flag"
	"fmt"
	"github.com/skr.io7803/pointerpw"
)

func main() {
	grpcAddress := flag.String("grpc", "localhost:8081", "address of GRPC server")
	restAddress := flag.String("rest", "localhost:8082", "address of REST server")
	flag.Parse()

	go func() {
		fmt.Printf("Running GRPC server at %s\n", *grpcAddress)
		if err := pointerpw.StartGRPC(*grpcAddress); err != nil {
			panic(err)
		}
		fmt.Printf("GRPC server stopped\n")
	}()

	go func() {
		fmt.Printf("Running REST server at %s\n", *restAddress)
		if err := pointerpw.StartREST(*restAddress, *grpcAddress); err != nil {
			panic(err)
		}
		fmt.Printf("REST server stopped\n")
	}()

	select {}
}
