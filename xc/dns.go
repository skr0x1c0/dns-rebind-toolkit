package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var namePattern = regexp.MustCompilePOSIX("^[A-Za-z0-9]+$")

type assignSubdomainRequest struct {
	Domain    string `json:"domain"`
	Ttl       uint32 `json:"ttl"`
	ReplaceOk bool   `json:"replaceOk"`
}

type DnsSubdomainManager struct {
	ctx context.Context
	req chan string
}

func NewDnsSubdomainManager(ctx context.Context) *DnsSubdomainManager {
	manager := &DnsSubdomainManager{
		ctx: ctx,
		req: make(chan string, 5),
	}
	go manager.start()
	return manager
}

func (d *DnsSubdomainManager) start() {
	Logger.Info("Start dns subdomain manager")
	defer Logger.Info("Stop dns subdomain manager")
	for {
		select {
		case name := <-d.req:
			Logger.Debugf("Assigning subdomain %s", name)
			if err := AssignDnsSubdomain(name); err != nil {
				Logger.Debug("cannot assign dns subdomain, error " + err.Error())
			}
			Logger.Debugf("Assigned subdomain %s", name)
		case <-d.ctx.Done():
			return
		}
	}
}

func (d *DnsSubdomainManager) Assign(name string) {
	d.req <- name
}

func AssignDnsSubdomain(name string) error {
	request := assignSubdomainRequest{
		Domain:    name,
		Ttl:       0,
		ReplaceOk: false,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("cannot marshal assign subdomain request, error: %v", err)
	}

	resp, err := http.Post(
		"http://api.pointer.pw/v1/ssrf/assign", "application/json", bytes.NewBuffer(body))

	if err != nil {
		return fmt.Errorf("cannot execute assign subdomain POST request, error: %v", err)
	}
	if resp.StatusCode == 200 {
		return nil
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read error response body")
	}

	return fmt.Errorf("cannot assign domain, error: %s", string(body))
}

type releaseSubdomainRequest struct {
	Domain string `json:"domain"`
}

func ReleaseDnsSubdomain(name string) error {
	request := releaseSubdomainRequest{
		Domain: name,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("cannot marshal release subdomain request, error %v", err)
	}

	resp, err := http.Post(
		"http://api.pointer.pw/v1/ssrf/release", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("cannot execute release subdomain request, error %v", err)
	}
	if resp.StatusCode == 200 {
		return nil
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read error response body")
	}
	return fmt.Errorf("cannot release subdomain, error: %s", string(body))
}

func BuildDnsDomain(name string) (string, error) {
	if len(name) != 12 || !namePattern.MatchString(name) {
		return "", fmt.Errorf("invalid name %s", name)
	}

	return name + ".dns.pointer.pw", nil
}
