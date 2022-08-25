package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g14com0/go-app/pkg/handler"
)

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)

	eh := handler.NewExpense(l)

	sm := http.NewServeMux()
	sm.Handle("/", eh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 100)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Graceful shutdown", sig)

	tc, cf := context.WithTimeout(context.Background(), 30*time.Second)
	defer cf()
	s.Shutdown(tc)
}
