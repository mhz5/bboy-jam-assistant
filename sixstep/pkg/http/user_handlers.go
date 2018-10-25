// user_handlers.go provides handlers for CRUD operations on users, as well as authentication/authorization.
package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"bboy-jam-assistant/sixstep/pkg/auth"
	"bboy-jam-assistant/sixstep/pkg/sessions"
	"github.com/gorilla/mux"

	"google.golang.org/appengine/log"
)

const (
	usernameKey = "username"
	passwordKey = "password"
)


func (rtr *Router) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.PostFormValue(usernameKey)
	password := r.PostFormValue(passwordKey)
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

	setCookie(w, usernameKey, u.Username)
	userJson, err := json.Marshal(u)
	w.Write(userJson)
}

func (rtr *Router) handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprintf("username cookie is not an int: %v", err), http.StatusBadRequest)
		return
	}
	if !auth.UserIsAuthorized(ctx, userId) {
		http.Error(w, "user not authorized to view data", http.StatusUnauthorized)
		return
	}

	u, err := rtr.userService.User(ctx, userId)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(u)
	w.Write(userJson)
}

func (rtr *Router) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.PostFormValue(usernameKey)
	password := r.PostFormValue(passwordKey)
	u, err := rtr.authService.Authenticate(ctx, username, password)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	s := sessions.NewUserSession(u)
	err = s.Save(w, r)
	if err != nil {
		log.Errorf(ctx, "couldn't create new session: %v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	setCookie(w, usernameKey, u.Username)
	userJson, err := json.Marshal(u)
	w.Write(userJson)
}
