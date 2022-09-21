package main

import (
	"context"
	"log"
	"microservices-go/product-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main()  {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	// create the handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// create a new ServeMux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// create a new server
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server without blocking the rest of operations
	go func ()  {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// create a channel and notify it if there is an interrupt or process kill
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until a signal is received
	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// gracefully shutdown the server with 30s max waiting time for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(ctx)
}