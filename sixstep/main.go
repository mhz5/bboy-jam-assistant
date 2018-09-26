// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"google.golang.org/appengine"
)

var router = mux.NewRouter()

// Register router to work with AppEngine.
func init() {
	http.Handle("/", router)
}

func main() {
	router.HandleFunc("/", handle)
	router.HandleFunc("/test", handleTest)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, bboy world!")
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a test")
}
