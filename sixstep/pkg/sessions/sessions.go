package sessions

import (
	"fmt"
	"net/http"

	"bboy-jam-assistant/sixstep/cmd/sixstep"

	"github.com/gorilla/sessions"
)


const (
	userSessionName = "user_session"

	// Keys to store values in session.
	userIdKey = "userId"
	usernameKey = "username"
)

// TODO: Think about 2 users on the same browser.
// TODO: Do you need to hash the sessions key?

// CookieStore saves sessions data in encrypted cookies to be stored on clients.
// Server need not persist sessions, but can decrypt sessions data instead.
// TODO: Handle setting the env variable "SESSION_KEY"
//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewCookieStore([]byte("SNTAHOEI"))


// NewUserSession creates and returns a new session representing a logged-in user.
func NewUserSession(u *sixstep.User) *Session {
	// TODO: What's the best way to nullify password field?
	s := New(userSessionName)
	s.sess.Values[userIdKey] = int64(u.Id)
	s.sess.Values[usernameKey] = u.Username
	return s
}

func New(name string) *Session {
	sess := sessions.NewSession(store, name)
	return &Session{sess}
}

// UserSession returns the user session associated with the provided request.
// Returns error if there was an issue looking up the session.
func UserSession(r *http.Request) (*Session, error) {
	return Get(userSessionName, r)
}

func Get(name string, r *http.Request) (*Session, error) {
	session, err := store.Get(r, name)
	if err != nil {
		return nil, err
	}
	if session.IsNew {
		return nil, fmt.Errorf("session doesn't exist")
	}
	return &Session{session}, nil
}

type Session struct {
	sess *sessions.Session
}

// Save saves the sessions by writing it to a cookie in the response.
func (s *Session) Save(w http.ResponseWriter, r *http.Request) error {
	return s.sess.Save(r, w)
}

// TODO: What's the best way to structure these accessors? A function per accessor?
// UserId returns the user id stored in the session.
func (s *Session) UserId() (int64, error) {
	// TODO: What is difference between type conversion and casting?
	// TODO: Fix error case
	userId := s.sess.Values[userIdKey]
	if userId == nil {
		return 0, fmt.Errorf("session does not contain userId")
	}
	return userId.(int64), nil
}

// UserId returns the username stored in the session.
func (s *Session) Username() (string, error) {
	// TODO: What is difference between type conversion and casting?
	username := s.sess.Values[usernameKey]
	if username == nil {
		return "", fmt.Errorf("session does not contain username")
	}
	return username.(string), nil
}

