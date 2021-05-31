
package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestEchoServer(t *testing.T) {

	 tests :=  []struct {
		name string
		data string
		method string
		ctype string
		statusExpected int
	}{

		{
			name: "Test Post Valid Json",
			data: `{"username": "xyz", "upload":"xyz"}`,
			method: "POST",
			ctype: "application/json",
			statusExpected: http.StatusOK,
		},
		{
			name: "Test Put Valid Json",
			data: `{"username": "xyz", "upload":"xyz"}`,
			method: "PUT",
			ctype: "application/json",
			statusExpected: http.StatusOK,
		},
		{
			name: "Test Get Json",
			data: "{}",
			method: "GET",
			ctype: "application/json",
			statusExpected: http.StatusMethodNotAllowed,
		},
		{
			name: "Test Post Valid Json but with echoed false",
			data: `{"username": "xyz", "upload":"xyz", "echoed":"false"}`,
			method: "POST",
			ctype: "application/json",
			statusExpected: http.StatusOK,
		},
		{
			name: "Test Post Valid Json but with echoed true",
			data: `{"username": "xyz", "upload":"xyz", "echoed":"true"}`,
			method: "POST",
			ctype: "application/json",
			statusExpected: http.StatusBadRequest,
		},
		{
			name: "Test Post Valid Json but with erroneous content-type",
			data: `{"username": "xyz", "upload":"xyz"}`,
			method: "POST",
			ctype: "application/xml",
			statusExpected: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, "http://127.0.0.1:3000/api/echo", strings.NewReader(test.data))
			req.Header.Set("Content-Type", test.ctype)
			w := httptest.NewRecorder()

			echoServer(w, req)

			resp := w.Result()


			if resp.StatusCode != test.statusExpected {
				t.Fatalf("Status Code: Expected %v but got %v", test.statusExpected, resp.StatusCode)
			}

			if test.statusExpected == http.StatusOK { // check content if we get a 200
				body, _ := ioutil.ReadAll(resp.Body)
				var received map[string]interface{}
				json.Unmarshal([]byte(body), &received)

				v, found := received["echoed"]
				if !found || v != "true" {
					t.Fatal("Echoed not found")
				}
			}
		})
	}
}
