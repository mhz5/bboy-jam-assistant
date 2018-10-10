package sixstep

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"bboy-jam-assistant/sixstep/src/session"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)


const (
	userKind    = "User"
	usernameKey = "username"
	passwordKey = "password"
)

var (
	allowedOrigin = os.Getenv("ALLOWED_ORIGIN")
)

type User struct {
	Username	 string
	PasswordHash string
}


func init() {
	router := mux.NewRouter()
	router.HandleFunc("/", handle)
	router.HandleFunc("/users", handleCreateUser).Methods("POST")
	router.HandleFunc("/login", handleLogin).Methods("POST")

	r := corsRouter(router)

	// Register router to work with AppEngine.
	http.Handle("/", r)
}

func Run() {
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, bboy world!")
}

// TODO: Migrate to protos.
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	user := &User {
		Username: r.PostFormValue(usernameKey),
		PasswordHash: r.PostFormValue(passwordKey),
	}
	key := datastore.NewIncompleteKey(ctx, userKind, nil)
	key, err := datastore.Put(ctx, key, user)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	id := strconv.FormatInt(key.IntID(), 10)
	s := session.New(id)
	err = s.Save(r, w)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

	// TODO: Write a meaningful response.
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	username := r.PostFormValue(usernameKey)
	password := r.PostFormValue(passwordKey)
	query := datastore.NewQuery(userKind).Filter("Username = ", username)

	results := query.Run(ctx)

	user := &User{}
	key, err := results.Next(user)
	if err == datastore.Done {
		// TODO: Write proper "user not found" response.
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	} else if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	if user.PasswordHash != password {
		http.Error(w, "Incorrect password", http.StatusInternalServerError)
		return
	}

	// Correct password
	id := strconv.FormatInt(key.IntID(), 10)
	s := session.New(id)
	err = s.Save(r, w)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	// TODO: Write a meaningful response.
}

func corsRouter(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	return c.Handler(h)
}
