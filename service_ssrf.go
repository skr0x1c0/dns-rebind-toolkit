package pointerpw

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/skr.io7803/pointerpw/pb"
)

type ssrfRegistryServer struct {
	srv pb.DnsRegistryServiceServer
}

func NewSSRFRegistryServer(srv pb.DnsRegistryServiceServer) pb.SSRFRegistryServiceServer {
	return &ssrfRegistryServer{
		srv: srv,
	}
}

func (s *ssrfRegistryServer) Assign(ctx context.Context, request *pb.SSRFAssignRequest) (*empty.Empty, error) {
	return s.srv.Assign(ctx, &pb.DnsAssignRequest{
		Domain:    request.Domain,
		Ip4:       []byte{127, 0, 0, 1},
		Ip6:       []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		Ttl:       request.Ttl,
		ReplaceOk: request.ReplaceOk,
	})
}

func (s *ssrfRegistryServer) Release(ctx context.Context, request *pb.SSRFReleaseRequest) (*empty.Empty, error) {
	return s.srv.Release(ctx, &pb.DnsReleaseRequest{
		Domain: request.Domain,
	})
}
