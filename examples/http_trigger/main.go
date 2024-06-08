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

package main

import (
	functions "github.com/mfcollins3/go-azure-functions/pkg/functions"
	"log"
)

func getContact(response *functions.Response, request functions.Request) error {
	httpRequest, err := request.HTTPRequest("request")
	if err != nil {
		return err
	}

	log.Printf("Received request: %s %s", httpRequest.Method, httpRequest.URL)
	return nil
}

func createContact(
	response *functions.Response,
	request functions.Request,
) error {
	httpRequest, err := request.HTTPRequest("request")
	if err != nil {
		return err
	}

	log.Printf("Received request: %s %s", httpRequest.Method, httpRequest.URL)
	return nil
}

func main() {
	functions.Function("CreateContact", createContact)
	functions.Function("GetContact", getContact)
	functions.Start()
}
