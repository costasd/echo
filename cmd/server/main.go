package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	//"fmt"
)

func main () {
	listen := "127.0.0.1:3000"

	go startHandler(listen)

	select {}
}

func startHandler(listen string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/echo", echoServer)
	http.ListenAndServe(listen, mux)
}


func echoServer (w http.ResponseWriter, r *http.Request) {

	switch r.Method {
		case http.MethodPost:
		case http.MethodPut:
		default:
			http.Error(w, "only POST/PUT are allowed", http.StatusMethodNotAllowed)
			return
	}
	//	Header = map[string][]string{

	//		"Accept-Encoding": {"gzip, deflate"},

	//		"Accept-Language": {"en-us"},

	//		"Foo": {"Bar", "two"},

	//	}

	if strings.ToLower(r.Header.Get("Content-Type")) != "application/json" {
		http.Error(w, "Content-Type must be set to application/json", http.StatusBadRequest)
		return
	}

	content, err := ioutil.ReadAll(r.Body)

	//fmt.Printf(string(content))
	if err != nil {
		http.Error(w, "error while reading body", http.StatusBadRequest)
		return
	}

	var received map[string]string // cast all as strings
	errj := json.Unmarshal([]byte(content), &received)
	if errj != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if received == nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	echoed, found := received["echoed"]
	if found && echoed == "true" {
		http.Error(w, "echoed cannot be set", http.StatusBadRequest)
	} else {
		received["echoed"] = "true"
		c, _ := json.Marshal(received)
		w.Header().Add("Content-Type", "application/json")
		w.Write(c)
	}
}
