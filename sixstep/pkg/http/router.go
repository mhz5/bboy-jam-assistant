package http

import (
	"fmt"
	"net/http"
	"os"


	"bboy-jam-assistant/sixstep/cmd/sixstep"
	"bboy-jam-assistant/sixstep/pkg/auth"
	"bboy-jam-assistant/sixstep/pkg/datastore"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

func (rtr *Router) Handle() {
	rtr.router.HandleFunc("/", handle)
	rtr.router.HandleFunc("/users", rtr.handleCreateUser).Methods("POST")
	rtr.router.HandleFunc("/login", rtr.handleLoginUser).Methods("POST")
	cRouter := rtr.corsRouter(rtr.router)

	// Register router to work with AppEngine.
	http.Handle("/", cRouter)
}

// handle handles the root path
func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, bboy world!")
}

func (rtr *Router) corsRouter(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	return c.Handler(h)
}
