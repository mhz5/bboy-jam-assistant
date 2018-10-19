package http

import (
	"fmt"
	"net/http"
	"os"

	"bboy-jam-assistant/sixstep/cmd/sixstep"
	"bboy-jam-assistant/sixstep/pkg/auth"
	"bboy-jam-assistant/sixstep/pkg/datastore"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type Router struct {
	userService sixstep.UserService
	authService sixstep.AuthService
	// TODO: Doesn't make much sense why Router has a field 'router'
	router  *mux.Router
}

var (
	_ sixstep.Router = &Router{}
	allowedOrigin = os.Getenv("ALLOWED_ORIGIN")
)

func NewRouter() *Router {
	return &Router{
		userService: datastore.NewUserService(),
		authService: auth.NewService(),
		router: mux.NewRouter(),
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

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, bboy world!")
}

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
	u := &sixstep.User {
		Username: username,
		PasswordHash: saltedPassword,
	}
	err = r.userService.CreateUser(ctx, u)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	sess := auth.NewSession(fmt.Sprint(u.Id))
	err = sess.Save(w, req)

	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	// TODO: Write a meaningful response.
	fmt.Fprint(w, "Success")
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

	sess := auth.NewSession(fmt.Sprint(u.Id))
	err = sess.Save(w, req)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	// TODO: Write a meaningful response.
	fmt.Fprintf(w, "Success")
}
