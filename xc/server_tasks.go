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

func NewTaskManager(ctx context.Context, config XCConfig) *TaskManager {
	manager := &TaskManager{
		xcQueue:  make(chan *Session),
		dnsQueue: make(chan *Session),
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
			case session := <-t.xcQueue:
				t.handleXCQueueTask(session)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (t *TaskManager) handleXCQueueTask(session *Session) {
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
}

func (t *TaskManager) handleDnsQueueTask(session *Session) {
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
}
