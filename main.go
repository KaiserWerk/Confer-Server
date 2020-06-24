package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	host	string
	authkey	string
)

func main() {
	flag.StringVar(&host, "h", "127.0.0.1:1663", "the host:port to use")
	flag.StringVar(&authkey, "k", "your-authkey", "the auth key the client should use")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", homepageHandler)
	router.HandleFunc("/file", authMid(fileGetHandler)).Methods("POST")
	router.HandleFunc("/file", authMid(fileSaveHandler)).Methods("PUT")

	// catch ctrl+c for graceful shutdown
	notify := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(notify, os.Interrupt)

	srv := &http.Server{
		Addr: 				host,
		Handler:            router,
		ReadTimeout:		1 * time.Second,
		WriteTimeout:       2 * time.Second,
		IdleTimeout:        30 * time.Second,
		ReadHeaderTimeout:  2 * time.Second,
	}

	go func() {
		<-notify
		fmt.Println("Initiating graceful shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		err := srv.Shutdown(ctx)
		if err != nil {
			panic("Could not gracefully shut down server: " + err.Error())
		}
		close(done)
	}()


	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("server error: %v\n", err.Error())
	}
	<-done
	fmt.Println("Server shutdown complete. Have a nice day!")
}

func authMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("X-AuthKey")
		if authkey == "" || h != authkey {
			w.WriteHeader(http.StatusForbidden)
			msg := httpResponse{
				Status:  "error",
				Code:    100,
				Message: "Access denied",
			}
			json.NewEncoder(w).Encode(&msg)
			return
		}
		next(w, r)
	}
}