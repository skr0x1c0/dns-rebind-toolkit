package pointerpw

import (
	"strings"
	"sync"
	"time"
)

type DnsQueryResult struct {
	QType uint16
	Time  time.Time
	Rcode int
}

type DnsQueryLog interface {
	GetAll(name string) []DnsQueryResult
	Put(name string, result DnsQueryResult)
}

type inMemoryDnsQueryLog struct {
	mu    sync.RWMutex
	store map[string][]DnsQueryResult
}

func NewInMemoryDnsQueryLog() DnsQueryLog {
	return &inMemoryDnsQueryLog{
		store: make(map[string][]DnsQueryResult),
	}
}

func (i *inMemoryDnsQueryLog) GetAll(name string) []DnsQueryResult {
	name = strings.ToLower(name)
	i.mu.RLock()
	defer i.mu.RUnlock()

	res, ok := i.store[name]
	if !ok {
		return []DnsQueryResult{}
	}
	return res
}

func (i *inMemoryDnsQueryLog) Put(name string, result DnsQueryResult) {
	name = strings.ToLower(name)
	i.mu.Lock()
	defer i.mu.Unlock()

	res, ok := i.store[name]
	if !ok {
		res = make([]DnsQueryResult, 0, 1)
	}
	i.store[name] = append(res, result)
}
