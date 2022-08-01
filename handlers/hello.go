package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct {
	l *log.Logger
}

func NewHelloHandler(l *log.Logger) *HelloHandler {
	return &HelloHandler{l}
}

func (handler *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.l.Println("We are serving hello")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, %s", d)
}
