package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodbyeHandler struct {
	l *log.Logger
}

func NewGoodbyeHandler(l *log.Logger) *GoodbyeHandler {
	return &GoodbyeHandler{l}
}

func (handler *GoodbyeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.l.Println("We are serving goodbye")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Goodbye, %s", d)
}
