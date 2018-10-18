package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// TODO: Think about 2 users on the same browser.
// TODO: Do you need to hash the sessions key?

// CookieStore saves sessions data in encrypted cookies to be stored on clients.
// Server need not persist sessions, but can decrypt sessions data instead.
// TODO: Handle setting the env variable "SESSION_KEY"
//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewCookieStore([]byte("SNTAHOEI"))

type Session struct {
	sess *sessions.Session
}

// NewSession creates and returns a new sessions associated with the provided name.
func NewSession(name string) *Session {
	sess := sessions.NewSession(store, name)
	return &Session{sess}
}

// GetSession returns the sessions associated with the provided name.
// If no sessions is associated with the name, returns nil.
func GetSession(name string, r *http.Request) (*Session, error) {
	session, err := store.Get(r, name)
	if err != nil {
		return nil, err
	}
	if session.IsNew {
		return nil, nil
	}
	return &Session{session}, nil
}

// Save saves the sessions by writing it to a cookie in the response.
func (s *Session) Save(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(s.sess)
	return s.sess.Save(r, w)
}

// TODO: Implement accessors to get specific values from the sessions.
