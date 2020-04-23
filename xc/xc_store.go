package main

import "sync"

type Session struct {
	Session   string `json:"session"`
	User      string `json:"user"`
	UserId    int    `json:"user_id"`
	ContextId int    `json:"context_id"`
	Locale    string `json:"locale"`
}

type XCStore interface {
	GetSession() (Session, error)
	SetSession(session Session)
}

type inMemoryXCStore struct {
	mu      sync.Mutex
	session *Session
}

func NewInMemoryXCStore() XCStore {
	return &inMemoryXCStore{
		mu:      sync.Mutex{},
		session: nil,
	}
}

func (i *inMemoryXCStore) GetSession() (Session, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.session == nil {
		return Session{}, ErrorInvalidSession
	}

	return *i.session, nil
}

func (i *inMemoryXCStore) SetSession(session Session) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.session = &session
}
