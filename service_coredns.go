package pointerpw

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"github.com/skr.io7803/pointerpw/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var ErrorUnsupportedQType = fmt.Errorf("unsupported QType")

type dnsServer struct {
	dnsStore DnsStore
	config   Config
	log      DnsQueryLog
}

func NewDnsServiceServer(config Config, dnsStore DnsStore, log DnsQueryLog) pb.DnsServiceServer {
	return &dnsServer{config: config, dnsStore: dnsStore, log: log}
}

func (d *dnsServer) respondA(domain string) (dns.RR, error) {
	record, err := d.dnsStore.Get(domain)
	if err != nil {
		return nil, err
	}

	if record.Ip4 == nil {
		d.log.Put(domain, DnsQueryResult{
			QType: dns.TypeA,
			Time:  time.Now(),
			Rcode: dns.RcodeServerFailure,
		})

		return nil, ErrorDnsRecordNotFound
	}

	d.log.Put(domain, DnsQueryResult{
		QType: dns.TypeA,
		Time:  time.Now(),
		Rcode: dns.RcodeSuccess,
	})

	return &dns.A{
		Hdr: dns.RR_Header{
			Name:   domain + "." + d.config.DnsRoot + ".",
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    record.Ttl,
		},
		A: record.Ip4,
	}, nil
}

func (d *dnsServer) respondAAAA(domain string) (dns.RR, error) {
	record, err := d.dnsStore.Get(domain)
	if err != nil {
		return nil, err
	}

	if record.Ip6 == nil {
		d.log.Put(domain, DnsQueryResult{
			QType: dns.TypeAAAA,
			Time:  time.Now(),
			Rcode: dns.RcodeServerFailure,
		})
		return nil, ErrorDnsRecordNotFound
	}

	d.log.Put(domain, DnsQueryResult{
		QType: dns.TypeAAAA,
		Time:  time.Now(),
		Rcode: dns.RcodeSuccess,
	})

	return &dns.AAAA{
		Hdr: dns.RR_Header{
			Name:   domain + "." + d.config.DnsRoot + ".",
			Rrtype: dns.TypeAAAA,
			Class:  dns.ClassINET,
			Ttl:    record.Ttl,
		},
		AAAA: record.Ip6,
	}, nil
}

func (d *dnsServer) respondQ(q dns.Question) (rr dns.RR, err error) {
	domain, err := d.parseQName(q.Name)
	if err != nil {
		return nil, err
	}

	Logger.Debug(domain)

	switch q.Qtype {
	case dns.TypeA:
		return d.respondA(domain)
	case dns.TypeAAAA:
		return d.respondAAAA(domain)
	default:
		return nil, ErrorUnsupportedQType
	}
}

func (d *dnsServer) parseQName(name string) (string, error) {

	root := "." + d.config.DnsRoot + "."

	invalidQNameErr := status.Error(codes.InvalidArgument, "invalid root")
	if len(name) > 253 {
		return "", invalidQNameErr
	}

	if len(name) < len(root)+1 {
		return "", invalidQNameErr
	}

	if name[len(name)-len(root):] != root {
		return "", invalidQNameErr
	}

	return name[:len(name)-len(root)], nil
}

func (d *dnsServer) Query(_ context.Context, packet *pb.DnsPacket) (*pb.DnsPacket, error) {
	request := &dns.Msg{}
	if err := request.Unpack(packet.GetMsg()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response := &dns.Msg{}
	response.SetReply(request)
	response.Authoritative = true

	for _, q := range response.Question {
		rr, err := d.respondQ(q)
		if err != nil {
			Logger.Debugf("Error handing question %s, %v", q, err)
			continue
		}
		response.Answer = append(response.Answer, rr)
	}

	if len(response.Answer) == 0 {
		response.Rcode = dns.RcodeServerFailure
	}

	out, err := response.Pack()
	if err != nil {
		Logger.Error(err)
		return nil, status.Error(codes.Internal, "cannot pack response")
	}

	return &pb.DnsPacket{Msg: out}, nil
}
