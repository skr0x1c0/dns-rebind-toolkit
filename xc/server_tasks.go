package main

import (
	"context"
	"time"
)

type TaskManager struct {
	xcQueue  chan *Session
	dnsQueue chan *Session
	xcClient *XCClient
}

type XCNegativeCacheResult struct {
	err  error
	time time.Time
}

type DnsRegistrationResult struct {
	err  error
	time time.Time
}

func NewTaskManager(ctx context.Context, config XCConfig) *TaskManager {
	manager := &TaskManager{
		xcQueue:  make(chan *Session, 8),
		dnsQueue: make(chan *Session, 8),
		xcClient: NewXCClient(config),
	}
	manager.start(ctx)
	return manager
}

func (t *TaskManager) start(ctx context.Context) {
	go func() {
		Logger.Info("start xc task manager")
		defer Logger.Info("stop xc task manager")

		for {
			select {
			case session := <-t.xcQueue:
				t.handleXCQueueTask(session)
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		Logger.Info("start dns task manager")
		defer Logger.Info("stop dns task manager")

		for {
			select {
			case session := <-t.dnsQueue:
				t.handleDnsQueueTask(session)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (t *TaskManager) handleXCQueueTask(session *Session) {
	Logger.Debug("handling xc task for ", session.id)
	defer Logger.Debug("done handling xc task for ", session.id)

	err := t.xcClient.TriggerNegativeCache(BuildDnsHost(session.dnsSdName))

	session.mu.Lock()
	defer session.mu.Unlock()
	if session.negCacheRes != nil {
		Logger.Warnf("possible duplicate xc queue task submission for session")
		return
	}
	result := XCNegativeCacheResult{}
	result.err = err
	result.time = time.Now()
	session.negCacheRes = &result

	t.dnsQueue <- session
}

func (t *TaskManager) handleDnsQueueTask(session *Session) {
	Logger.Debug("handling dns task for ", session.id)
	defer Logger.Debug("done handling dns task for ", session.id)

	err := AssignDnsSubdomain(session.dnsSdName)

	session.mu.Lock()
	defer session.mu.Unlock()
	if session.dnsRegRes != nil {
		Logger.Warnf("possible duplicate dns registration task submission for session")
		return
	}
	result := DnsRegistrationResult{}
	result.err = err
	result.time = time.Now()
	session.dnsRegRes = &result
}

func (t *TaskManager) Submit(session *Session) {
	t.xcQueue <- session
}
