// user.go provides handlers for CRUD operations on users, as well as authentication/authorization.
package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bboy-jam-assistant/sixstep/pkg/auth"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	userSession = "user_session"
)

func (r *Router) handleCreateUser(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	username := req.PostFormValue("username")
	password := req.PostFormValue("password")
	saltedPassword, err := auth.GenerateSaltedPassword(password)
	// TODO: How to remove redundancy in error handling?
	if err != nil {
		http.Error(w, fmt.Sprintf("bad password: %v", err), http.StatusInternalServerError)
		return
	}

	u, err := r.userService.CreateUser(ctx, username, saltedPassword)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	s := auth.NewSession(userSession)
	err = s.Save(w, req)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(u)
	w.Write(userJson)
}

func (r *Router) handleLogin(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	username := req.PostFormValue("username")
	password := req.PostFormValue("password")
	u, err := r.authService.Authenticate(ctx, username, password)

	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	s := auth.NewSession(userSession)
	err = s.Save(w, req)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	userJson, err := json.Marshal(u)
	w.Write(userJson)
}
