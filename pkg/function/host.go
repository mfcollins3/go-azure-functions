// Copyright 2024 Michael F. Collins, III
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of thi software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDER BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.
//

package functions

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func Function(
	pattern string,
	handler func(*Response, Request),
) {
	http.HandleFunc(
		pattern,
		func(w http.ResponseWriter, r *http.Request) {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Println(string(body))

			var request Request
			if err = json.Unmarshal(body, &request); err != nil {
				log.Printf("ERROR: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			response := NewResponse()
			handler(response, request)

			bytes, err := json.Marshal(response)
			if err != nil {
				log.Printf("ERROR: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(bytes); err != nil {
				log.Printf("ERROR: %v", err)
			}
		},
	)
}

// Start starts the custom worker that listens for incoming Azure
// Function requests from the Azure Functions host.
//
// Start listens on the port specified by the
// FUNCTIONS_CUSTOMHANDLER_PORT environment variable. If the environment
// variable is not set, ListenAndServe logs a fatal error.
func Start(handler http.Handler) {
	var listenAddress string
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddress = ":" + val
	} else {
		log.Fatal("the FUNCTIONS_CUSTOMHANDLER_PORT environment variable is not set")
	}

	log.Fatal(http.ListenAndServe(listenAddress, handler))
}
