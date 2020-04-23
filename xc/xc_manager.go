package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	ErrorJobAlreadySubmitted = fmt.Errorf("job already submitted")
)

type XCNegativeCacheResult struct {
	done bool
	err  error
	time time.Time
}

type XCManager struct {
	mu        sync.Mutex
	store     map[string]XCNegativeCacheResult
	queue     chan string
	client    *XCClient
	dsManager *DnsSubdomainManager
}

func NewXCManager(ctx context.Context, client *XCClient) *XCManager {
	if client == nil {
		panic("expected valid client")
	}

	manager := &XCManager{
		store:     make(map[string]XCNegativeCacheResult),
		queue:     make(chan string, 5),
		client:    client,
		dsManager: NewDnsSubdomainManager(ctx),
	}

	go manager.start(ctx)

	return manager
}

func (x *XCManager) start(ctx context.Context) {
	fmt.Println("Start XCManager")
	defer fmt.Println("Stop XCManager")
	for {
		select {
		case name := <-x.queue:
			domain, err := BuildDnsDomain(name)
			AssertOk(err)
			err = x.client.TriggerNegativeCache(domain)
			x.update(name, err)
			x.dsManager.Assign(name)
		case <-ctx.Done():
			return
		}
	}
}

func (x *XCManager) update(name string, err error) {
	x.mu.Lock()
	defer x.mu.Unlock()

	fmt.Printf("Completed negative cache trigger for %s\n", name)
	x.store[name] = XCNegativeCacheResult{
		done: true,
		err:  err,
		time: time.Now(),
	}
}

func (x *XCManager) Submit(name string) error {
	x.mu.Lock()
	defer x.mu.Unlock()

	if _, ok := x.store[name]; ok {
		return ErrorJobAlreadySubmitted
	}

	x.store[name] = XCNegativeCacheResult{
		done: false,
		err:  nil,
		time: time.Unix(0, 0),
	}

	x.queue <- name
	return nil
}

func (x *XCManager) Status(name string) (XCNegativeCacheResult, bool) {
	x.mu.Lock()
	defer x.mu.Unlock()
	v, ok := x.store[name]
	return v, ok
}
