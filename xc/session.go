package main

import (
	"net/url"
	"sync"
	"time"
)

type Scheme string

const (
	SchemeHttp  Scheme = "http"
	SchemeHttps        = "https"
)

type Session struct {
	id            string
	dnsSdName     string
	payloadSize   uint64
	sleepDuration time.Duration
	targetUrl     *url.URL
	targetPort    uint16
	targetScheme  Scheme

	mu          sync.Mutex
	negCacheRes *XCNegativeCacheResult
	dnsRegRes   *DnsRegistrationResult

	startTime *time.Time
}
