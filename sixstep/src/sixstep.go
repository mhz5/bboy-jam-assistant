// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sixstep

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

var (
	router = mux.NewRouter()
	AllowedOrigin = os.Getenv("ALLOWED_ORIGIN")
)

type User struct {
	Username	 string
	PasswordHash string
}

// Register router to work with AppEngine.
func init() {
	http.Handle("/", router)
}

func Run() {
	router.HandleFunc("/", injectCors(handle))
	router.HandleFunc("/test", injectCors(handleTest))
	router.HandleFunc("/users", injectCors(handleCreateUser)).Methods("POST")
	router.HandleFunc("/users", injectCors(handleCreateUserOption)).Methods("OPTIONS")

	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	// ctx := appengine.NewContext(r)
	// Example of getting env variable.
	// log.Infof(ctx, os.Getenv("ALLOWED_ORIGIN"))
	fmt.Fprintln(w, "Hello, bboy world!")
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a test")
}

func handleCreateUserOption(w http.ResponseWriter, r *http.Request) {
	// TODO: Figure out correct way to handle preflight CORS request.
	w.WriteHeader(http.StatusOK)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	key := datastore.NewIncompleteKey(ctx, "User", nil)

	user := &User {
		Username: r.PostFormValue("username"),
		PasswordHash: r.PostFormValue("password"),
	}
	_, err := datastore.Put(ctx, key, user)

	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

	resp := struct {
		UserId string
		SessionKey string
	}{
		UserId: "userId",
		SessionKey: "sessionKey",
	}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(b))
}

func injectCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedOrigin)
		next(w, r)
	}
}
