package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Goodbye handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodBye creates a new goodbye handler with the given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHTTP implements the go http.Handlers interface
// https://pkg.go.dev/net/http#Handler
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle Goodbye request")

	fmt.Fprintf(rw, "Goodbye")
}