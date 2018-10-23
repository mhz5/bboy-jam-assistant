// user.go provides handlers for CRUD operations on users, as well as authentication/authorization.
package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bboy-jam-assistant/sixstep/pkg/auth"
	"bboy-jam-assistant/sixstep/pkg/sessions"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)


func (rtr *Router) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	saltedPassword, err := auth.GenerateSaltedPassword(password)
	// TODO: How to remove redundancy in error handling?
	if err != nil {
		http.Error(w, fmt.Sprintf("bad password: %v", err), http.StatusInternalServerError)
		return
	}

	u, err := rtr.userService.CreateUser(ctx, username, saltedPassword)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	s := sessions.NewUserSession(u)
	err = s.Save(w, r)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(u)
	w.Write(userJson)
}

func (rtr *Router) handleGetUser(w http.ResponseWriter, r *http.Request) {
	r.Context()
}

func (rtr *Router) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	u, err := rtr.authService.Authenticate(ctx, username, password)

	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	s := sessions.NewUserSession(u)
	err = s.Save(w, r)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	userJson, err := json.Marshal(u)
	w.Write(userJson)
}
