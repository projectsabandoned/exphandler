package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hi struct {
	l *log.Logger
}

func NewHi(l *log.Logger) *Hi {
	return &Hi{l}
}

func (h *Hi) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	fmt.Fprintf(rw, "Hi %s", data)

	if err != nil {
		http.Error(rw, "Error: ", http.StatusBadRequest)
		return
	}
}
