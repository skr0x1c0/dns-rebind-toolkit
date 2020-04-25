package main

import (
	"net/url"
	"time"
)

type Scheme string

const (
	SchemeHttp  Scheme = "http"
	SchemeHttps        = "https"
)

type Session struct {
	dnsSdName     string
	payloadSize   uint64
	sleepDuration time.Duration
	targetUrl     *url.URL
	targetPort    uint16
	targetScheme  Scheme

	negCacheRes *XCNegativeCacheResult
	dnsRegRes   *DnsRegistrationResult

	startTime *time.Time
}

type DnsRegistrationResult struct {
	err  error
	time time.Time
}
