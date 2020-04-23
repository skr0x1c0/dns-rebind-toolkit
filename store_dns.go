package pointerpw

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var ErrorDnsRecordExist = fmt.Errorf("dns record already exist")
var ErrorDnsRecordNotFound = fmt.Errorf("dns record not found")

type DnsStore interface {
	Get(domain string) (Record, error)
	Set(domain string, record Record, replaceOk bool) error
	Remove(domain string) error
}

type Record struct {
	Ip4 net.IP
	Ip6 net.IP
	Ttl uint32
}

type inMemoryDnsStore struct {
	mu    sync.RWMutex
	store map[string]Record
}

func NewInMemoryDnsStore() DnsStore {
	return &inMemoryDnsStore{
		mu:    sync.RWMutex{},
		store: make(map[string]Record),
	}
}

func (i *inMemoryDnsStore) Get(domain string) (record Record, err error) {
	domain = i.prepName(domain)
	i.mu.RLock()
	defer i.mu.RUnlock()

	var ok bool
	if record, ok = i.store[domain]; !ok {
		err = ErrorDnsRecordNotFound
	}
	return
}

func (i *inMemoryDnsStore) Set(domain string, record Record, replaceOk bool) error {
	domain = i.prepName(domain)
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, ok := i.store[domain]; ok && !replaceOk {
		return ErrorDnsRecordExist
	}
	i.store[domain] = record
	return nil
}

func (i *inMemoryDnsStore) Remove(domain string) error {
	domain = i.prepName(domain)
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, ok := i.store[domain]; !ok {
		return ErrorDnsRecordNotFound
	}
	delete(i.store, domain)
	return nil
}

func (i *inMemoryDnsStore) prepName(name string) string {
	return strings.ToLower(name)
}
