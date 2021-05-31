package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
)

func main() {
	listen := "127.0.0.1:3000"

	go startHandler(listen)

	select {}
}

func startHandler(listen string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/echo", echoServer)
	http.ListenAndServe(listen, mux)
}

func echoServer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
	case http.MethodPut:
	default:
		http.Error(w, "only POST/PUT are allowed", http.StatusMethodNotAllowed)
		log.Printf("%v [%v] %v %v", http.StatusMethodNotAllowed, r.RemoteAddr, "/api/echo", "Error: only POST/PUT are allowed")
		return
	}

	if strings.ToLower(r.Header.Get("Content-Type")) != "application/json" {
		http.Error(w, "Content-Type must be set to application/json", http.StatusBadRequest)
		log.Printf("%v [%v] %v %v", http.StatusBadRequest, r.RemoteAddr, "/api/echo", "Error: Content-Type must be set to application/json")
		return
	}

	content, err := ioutil.ReadAll(r.Body)

	//fmt.Printf(string(content))
	if err != nil {
		http.Error(w, "error while reading body", http.StatusBadRequest)
		log.Printf("%v [%v] %v %v", http.StatusBadRequest, r.RemoteAddr, "/api/echo", "Error: while reading body")
		return
	}

	var received map[string]string // cast all as strings
	errj := json.Unmarshal([]byte(content), &received)
	if errj != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		log.Printf("%v [%v] %v %v", http.StatusBadRequest, r.RemoteAddr, "/api/echo", "Error: invalid json")
		return
	}

	if received == nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		log.Printf("%v [%v] %v %v", http.StatusBadRequest, r.RemoteAddr, "/api/echo", "Error: invalid json")
		return
	}

	echoed, found := received["echoed"]
	if found && echoed == "true" {
		http.Error(w, "echoed cannot be set", http.StatusBadRequest)
		log.Printf("%v [%v] %v %v", http.StatusBadRequest, r.RemoteAddr, "/api/echo", "Error: echoed cannot be set to true")
	} else {
		received["echoed"] = "true"
		c, _ := json.Marshal(received)
		w.Header().Add("Content-Type", "application/json")
		w.Write(c)
		log.Printf("%v [%v] %v %v", http.StatusOK, r.RemoteAddr, "/api/echo", "OK: echoed back")
	}
}
