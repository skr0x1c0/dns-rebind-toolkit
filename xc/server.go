package main

import (
	"context"
	"flag"
	"net/http"
	"time"
)

const (
	MaxResponseTime = 1500 * time.Millisecond
	KB              = 1024
	MB              = 1024 * KB
	MaxPayloadSize  = 50 * MB
)

func main() {
	address := flag.String("address", "172.16.146.1:8085", "server address")
	xcRoot := flag.String("xcRoot", "http://172.16.66.130", "oxc server root url")
	xcUser := flag.String("xcUser", "testuser", "oxc username")
	xcPassword := flag.String("xcPassword", "secret", "oxc user password")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := NewServerHandler(ctx, XCConfig{
		Root:     *xcRoot,
		Username: *xcUser,
		Password: *xcPassword,
	})

	server := &http.Server{
		Addr:         *address,
		Handler:      makeHttpHandler(handler),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	Logger.Info("Start server at ", *address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func makeHttpHandler(handler *ServerHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Infow(r.Method, " ", r.URL.Path, " ", r.RemoteAddr, " ", r.UserAgent())

		var err error
		if r.URL.Path == "/new" && r.Method == http.MethodGet {
			err = handler.HandleNewSession(w, r)
		} else if r.URL.Path == "/redirect" && r.Method == http.MethodGet {
			err = handler.HandleRedirect(w, r)
		} else {
			http.NotFound(w, r)
			return
		}

		switch e := err.(type) {
		case nil:
			return
		case *BadRequest:
			Logger.Debug("got bad request, msg: %s", e.msg)
			http.Error(w, e.msg, http.StatusBadRequest)
		case *InternalError:
			Logger.Errorf("internal error, msg: %s, err: %+v", e.msg, e.detail)
			http.Error(w, e.msg, http.StatusInternalServerError)
		default:
			Logger.Errorf("unexpected error, err: %+v", e)
			http.Error(w, "unexptected error", http.StatusInternalServerError)
		}
	})
}
