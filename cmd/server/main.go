package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
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

	c, _ := json.Marshal(received)
	w.Write(c)
}
