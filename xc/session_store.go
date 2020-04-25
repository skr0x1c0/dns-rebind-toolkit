package main

import (
	"fmt"
	"sync"
)

var (
	ErrorSessionExist      = fmt.Errorf("session already exist")
	ErrorSessionDoNotExist = fmt.Errorf("session do not exist")
)

type SessionStore interface {
	Get(key string) (*Session, error)
	Set(key string, val *Session) error
	Remove(key string) error
}

type inMemorySessionStore struct {
	mu    sync.Mutex
	store map[string]*Session
}

func NewInMemorySessionStore() SessionStore {
	return &inMemorySessionStore{
		store: make(map[string]*Session),
	}
}

func (i *inMemorySessionStore) Get(key string) (*Session, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if session, ok := i.store[key]; ok {
		return session, nil
	}
	return nil, ErrorSessionDoNotExist
}

func (i *inMemorySessionStore) Set(key string, val *Session) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, ok := i.store[key]; ok {
		return ErrorSessionExist
	}
	i.store[key] = val
	return nil
}

func (i *inMemorySessionStore) Remove(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, ok := i.store[key]; !ok {
		return ErrorSessionDoNotExist
	}
	delete(i.store, key)
	return nil
}
