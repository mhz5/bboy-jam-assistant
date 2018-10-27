package http

import (
	"fmt"
	"net/http"

	"bboy-jam-assistant/sixstep/cmd/sixstep"
	"bboy-jam-assistant/sixstep/pkg/auth"
	"bboy-jam-assistant/sixstep/pkg/datastore"

	"github.com/gorilla/mux"
)

type Router struct {
	userService sixstep.UserService
	authService sixstep.AuthService
	// TODO: Investigate using mux.Router.middlewares
	*mux.Router
}

var (
	_ sixstep.Router = &Router{}
)

func NewRouter() *Router {
	return &Router{
		datastore.NewUserService(),
		auth.NewService(),
		mux.NewRouter(),
	}
}

// We name receiver `rtr` because `r` is reserved for *http.Request in handlers.
func (rtr *Router) Handle() {
	rtr.HandleFunc("/", handle)
	rtr.HandleFunc("/users/{userId}", authorize(rtr.handleGetUser)).Methods("GET")
	rtr.HandleFunc("/users", rtr.handleCreateUser).Methods("POST")
	rtr.HandleFunc("/login", rtr.handleLoginUser).Methods("POST")
	router := warmup(corsRouter(appengineCtxRouter(rtr)))

	// Register router to work with AppEngine.
	http.Handle("/", router)
}

// handle handles the root path
func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, bboy world!")
}
