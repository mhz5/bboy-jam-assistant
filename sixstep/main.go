// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"net/http"
	"os"
)

const (
	AllowedOriginEnvKey = "ALLOWED_ORIGIN"
)

var router = mux.NewRouter()

// Register router to work with AppEngine.
func init() {
	http.Handle("/", router)
}

func main() {
	router.HandleFunc("/", injectCors(handle))
	router.HandleFunc("/test", injectCors(handleTest))
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Example of getting env variable.
	//ctx := appengine.NewContext(r)
	//log.Infof(ctx, os.Getenv("CLIENT_URL"))

	fmt.Fprintln(w, "Hello, bboy world!")
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a test")
}

func injectCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := os.Getenv(AllowedOriginEnvKey)
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		next(w, r)
	}
}
