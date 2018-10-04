package session

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// CookieStore saves session data in encrypted cookies to be stored on clients.
// Server need not persist session, but can decrypt session data instead.
// TODO: Handle setting the env variable "SESSION_KEY"
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type Session struct {
	sess *sessions.Session
}

// New creates and returns a new session associated with the provided name.
func New(name string) *Session {
	sess := sessions.NewSession(store, name)
	return &Session{sess}
}

// Get returns the session associated with the provided name.
// If no session is associated with the name, returns nil.
func Get(name string, r *http.Request) (*Session, error) {
	session, err := store.Get(r, name)
	if err != nil {
		return nil, err
	}
	if session.IsNew {
		return nil, nil
	}
	return &Session{session}, nil
}

// Save saves the session by writing it to a cookie in the response.
func (s *Session) Save(w http.ResponseWriter) {
	s.sess.Save(nil, w)
}

// TODO: Implement accessors to get specific values from the session.
