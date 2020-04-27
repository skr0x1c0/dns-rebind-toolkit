package pointerpw

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/skr.io7803/pointerpw/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

var domainNameRegex = regexp.MustCompilePOSIX("^[A-Za-z0-9]+$")

type dnsManagerService struct {
	config   Config
	dnsStore DnsStore
	log      DnsQueryLog
}

func NewDnsRegistryServiceServer(config Config, dnsStore DnsStore, log DnsQueryLog) pb.DnsRegistryServiceServer {
	return &dnsManagerService{config: config, dnsStore: dnsStore, log: log}
}

func (d *dnsManagerService) validateDomain(domain string) error {
	if !domainNameRegex.MatchString(domain) {
		return status.Error(codes.InvalidArgument, "invalid domain name")
	}
	maxLength := 253 - len(d.config.DnsRoot) - 2
	if len(domain) > maxLength {
		return status.Error(codes.InvalidArgument,
			fmt.Sprintf("domain name length should be less than %d", maxLength))
	}
	return nil
}

func (d *dnsManagerService) validateIp4(data []byte) error {
	if data != nil && len(data) != 4 {
		return status.Error(codes.InvalidArgument, "invalid Ip4")
	}
	return nil
}

func (d *dnsManagerService) validateIp6(data []byte) error {
	if data != nil && len(data) != 16 {
		return status.Error(codes.InvalidArgument, "invalid Ip6")
	}
	return nil
}

func (d *dnsManagerService) translateDnsStoreError(err error) error {
	if err == nil {
		panic("valid error expected")
	}
	if err == ErrorDnsRecordNotFound {
		return status.Error(codes.NotFound, "dns record not found")
	}
	if err == ErrorDnsRecordExist {
		return status.Error(codes.AlreadyExists, "dns record already exist")
	}
	return status.Error(codes.Internal, "unknown dns store error")
}

func (d *dnsManagerService) Assign(_ context.Context, request *pb.DnsAssignRequest) (*empty.Empty, error) {
	if err := d.validateDomain(request.Domain); err != nil {
		return nil, err
	}
	if err := d.validateIp4(request.Ip4); err != nil {
		return nil, err
	}
	if err := d.validateIp6(request.Ip6); err != nil {
		return nil, err
	}

	if err := d.dnsStore.Set(request.Domain, Record{
		Ip4: request.Ip4,
		Ip6: request.Ip6,
		Ttl: request.Ttl,
	}, request.ReplaceOk); err != nil {
		return nil, d.translateDnsStoreError(err)
	}

	return &empty.Empty{}, nil
}

func (d *dnsManagerService) Release(_ context.Context, request *pb.DnsReleaseRequest) (*empty.Empty, error) {
	if err := d.validateDomain(request.Domain); err != nil {
		return nil, err
	}

	if err := d.dnsStore.Remove(request.Domain); err != nil {
		return nil, d.translateDnsStoreError(err)
	}

	return &empty.Empty{}, nil
}

func (d *dnsManagerService) GetLog(_ context.Context, request *pb.DnsGetLogRequest) (*pb.DnsGetLogResponse, error) {
	name := request.Domain
	logs := d.log.GetAll(name)

	result := make([]*pb.DnsGetLogResponse_DnsLog, len(logs))
	for idx, v := range logs {
		result[idx] = &pb.DnsGetLogResponse_DnsLog{}
		result[idx].QType = uint32(v.QType)
		result[idx].RCode = int32(v.Rcode)
		result[idx].Timestamp = v.Time.Unix()
	}

	return &pb.DnsGetLogResponse{Log: result}, nil
}
