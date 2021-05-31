
package main

import (
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
		statusExpected int
	}{

		{
			name: "Test Post Valid Json",
			data: `{"username": "xyz", "upload":"xyz"}`,
			method: "POST",
			statusExpected: http.StatusOK,
		},
		{
			name: "Test Put Valid Json",
			data: `{"username": "xyz", "upload":"xyz"}`,
			method: "PUT",
			statusExpected: http.StatusOK,
		},
		{
			name: "Test Get Json",
			data: "{}",
			method: "GET",
			statusExpected: http.StatusOK,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, "http://127.0.0.1:3000/api/echo", strings.NewReader(test.data))
			w := httptest.NewRecorder()

			echoServer(w, req)

			resp := w.Result()


			if resp.StatusCode != test.statusExpected {
				t.Fatalf("Status Code: Expected %v but got %v", test.statusExpected, resp.StatusCode)
			}
		})
	}
}
