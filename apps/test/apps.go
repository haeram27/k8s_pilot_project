package main

import (
	"fmt"
	"net/http"
	"os"
)

type appHandler struct {
	http.Handler
}

func (h *appHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	buff := fmt.Sprintf("you've hit %s\n", hostname)

	res.Write([]byte(buff))
}

func main() {
	http.Handle("/", new(appHandler))

	http.ListenAndServe(":8080", nil)
}
