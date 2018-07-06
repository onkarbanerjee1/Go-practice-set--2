package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	//the entries map is used as the store to record the entries
	entries *map[string]string

	//for logging purpose
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

//intiate the InfoLogger and ErrorLogger
func InitLogger(infoHandle io.Writer, errorHandle io.Writer) {
	InfoLogger = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Main function to support Get(to retrieve), Put(to create and update any record), Delete(to delete record) and
// Count(to get count of key matched by prefixes.
func main() {

	InitLogger(os.Stdout, os.Stderr)
	//use th mux package to create routes and assign handlers to each of them accoridngly along with the HTTP methods
	h := http.NewServeMux()
	data := make(map[string]string)
	entries = &data
	// Routes consist of a path and a handler function.
	h.HandleFunc("/redis/entries/", EntriesEndPointHandler)

	h.HandleFunc("/redis/count", CountEndPoint)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8765", h))
}

//function to match the endpoint handlers based on the HTTP method
func EntriesEndPointHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetEndPoint(w, r)

	case "DELETE":
		DelEndPoint(w, r)

	case "POST":
		PutEndPoint(w, r)

	default:
		DefaultBadRequest(w)
	}
}

//GetEndPoint is used here to read the key value from the path params in the url
//and corresponding value returned else not found error returned
func GetEndPoint(w http.ResponseWriter, r *http.Request) {

	params := strings.Split(strings.TrimPrefix(r.URL.Path, "/redis/entries/"), "/")
	key := params[0]
	if val, ok := (*entries)[key]; ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", val)
		InfoLogger.Printf("Returning value -> %s for resource -> %s", val, key)
	} else {
		http.NotFound(w, r)
		ErrorLogger.Printf("Resource -> %s not found", key)
	}

}

// PutEndPoint for creating an entry with a kay and a val, retrieved from the form data
func PutEndPoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		key := r.FormValue("key")
		value := r.FormValue("value")
		if len(strings.TrimSpace(key)) > 0 {
			(*entries)[key] = value
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Record entered")
			InfoLogger.Printf("Record entered with key -> %s and value -> %s", key, value)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Empty data for -> key , need non-empty data for this field")
			ErrorLogger.Println("Bad request received with empty data for -> key")
		}
	} else {
		DefaultBadRequest(w)
	}
}

// DelEndPoint to check if key retrieved in path param of url exists or not and
// remove the record corresponding to the key if exists
func DelEndPoint(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(strings.TrimPrefix(r.URL.Path, "/redis/entries/"), "/")
	key := params[0]
	if val, ok := (*entries)[key]; ok {
		delete(*entries, key)
		fmt.Fprint(w, "Record deleted")
		InfoLogger.Printf("Record deleted with key -> %s and val -> %s", key, val)
	} else {
		fmt.Fprintf(w, "Key -> %s does not exist", key)
		ErrorLogger.Printf("Delete request for Key -> %s which does not exist", key)
	}

}

//CountEndPoint - check if url contains any query param "key", if so get the
//count accordingly by prefix match else return bad request received
func CountEndPoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if search, ok := r.URL.Query()["key"]; ok {
			w.WriteHeader(http.StatusOK)
			var count int
			for k := range *entries {
				if strings.HasPrefix(k, search[0]) {
					count++
				}
			}
			fmt.Fprintf(w, "%x", count)
			InfoLogger.Printf("Returning count for %s = %d", search[0], count)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "To retrieve count use query param \"key=\" ")
			ErrorLogger.Println("Bad request received for count without query param = \"key=\"")
		}
	} else {
		DefaultBadRequest(w)
	}

}

//function to write to response in case of bad request(ie wrong URL and/or method)
func DefaultBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Please check URL and method")
}
