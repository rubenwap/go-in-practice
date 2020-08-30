package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"github.com/braintree/manners"
)

func main() {
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	manners.ListenAndServe(":8080", handler)
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

// Note that the manners package only works for http connections. 
// If you are dealing with TCP, this won't work
func listenForShutdown(ch <- chan os.Signal) {
	<- ch
	fmt.Println("Received shutdown signal")
	manners.Close()
	fmt.Println("Gracefully closed")
}