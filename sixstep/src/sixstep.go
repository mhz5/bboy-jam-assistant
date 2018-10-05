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

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	user := &User {
		Username: r.PostFormValue("username"),
		PasswordHash: r.PostFormValue("password"),
	}
	key := datastore.NewIncompleteKey(ctx, "User", nil)
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
