package http

import (
	"fmt"
	"net/http"
	"os"

	"bboy-jam-assistant/sixstep/cmd/sixstep"
	"bboy-jam-assistant/sixstep/pkg/auth"
	"bboy-jam-assistant/sixstep/pkg/datastore"

	"github.com/gorilla/mux"
)

type Router struct {
	userService sixstep.UserService
	authService sixstep.AuthService
	// TODO: Doesn't make much sense why Router has a field 'router'. Read about interfaces.
	router *mux.Router
}

var (
	_ sixstep.Router = &Router{}
	allowedOrigin = os.Getenv("ALLOWED_ORIGIN")
)

func NewRouter() *Router {
	return &Router{
		userService: datastore.NewUserService(),
		authService: auth.NewService(),
		router:      mux.NewRouter(),
	}
}

func (r *Router) Handle() {
	r.router.HandleFunc("/", handle)
	r.router.HandleFunc("/users", r.handleCreateUser).Methods("POST")
	r.router.HandleFunc("/login", r.handleLogin).Methods("POST")
	cRouter := r.corsRouter(r.router)

	// Register router to work with AppEngine.
	http.Handle("/", cRouter)
}

// handle handles the root path
func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, bboy world!")
}
