package main

import (
	"bytes"
	"crypto/md5"
	"fmt"

	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Routes consist of a path and a handler function.

	mux.HandleFunc("/md5sum", ReqHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%x", md5sum(r.FormValue("what")))

	default:

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Not found any post request")

	}

}

// function to calculate the md5sum and return the same
func md5sum(data string) string {

	h := md5.New()
	if _, err := io.Copy(h, bytes.NewBufferString(data)); err != nil {
		log.Fatal(err)
	}
	b := h.Sum(nil)
	fmt.Printf("%x", b)
	return string(b[:])
}
