package pointerpw

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/skr.io7803/pointerpw/pb"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func initStore(store DnsStore) {
	err := store.Set("local", Record{
		Ip4: net.IPv4(127, 0, 0, 1),
		Ip6: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		Ttl: 0,
	}, false)
	if err != nil {
		panic(err)
	}
}

func StartGRPC(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("cannot listen on address %s, error %v", address, err)
	}

	config := Config{
		DnsRoot: "dns.pointer.pw",
	}
	dnsStore := NewInMemoryDnsStore()
	initStore(dnsStore)

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpczap.StreamServerInterceptor(Logger.Desugar()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpczap.UnaryServerInterceptor(Logger.Desugar()),
		)))
	pb.RegisterDnsServiceServer(server, NewDnsServiceServer(config, dnsStore))
	pb.RegisterDnsRegistryServiceServer(server, NewDnsRegistryServiceServer(config, dnsStore))
	pb.RegisterSSRFRegistryServiceServer(server, NewSSRFRegistryServer(
		NewDnsRegistryServiceServer(config, dnsStore)))

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("cannot start grpc server, error %v", err)
	}

	return nil
}

func StartREST(address string, grpcAddress string) error {
	ctx := context.Background()

	client, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	if err := pb.RegisterSSRFRegistryServiceHandler(ctx, mux, client); err != nil {
		panic(err)
	}

	return http.ListenAndServe(address, mux)
}
