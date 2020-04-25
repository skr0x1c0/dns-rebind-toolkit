package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DnsSubdomainLength int = 12
)

type assignSubdomainRequest struct {
	Domain    string `json:"domain"`
	Ttl       uint32 `json:"ttl"`
	ReplaceOk bool   `json:"replaceOk"`
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

func BuildDnsHost(name string) string {
	return name + ".dns.pointer.pw"
}
