package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)

		fmt.Fprintf(rw, "%s", data)

		if err != nil {
			http.Error(rw, "Error: ", http.StatusBadRequest)
			return
		}
	})

	http.ListenAndServe(":9090", nil)
}
