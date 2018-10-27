package sessions

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"bboy-jam-assistant/sixstep/cmd/sixstep"

	"cloud.google.com/go/storage"
	"github.com/gorilla/sessions"
)


const (
	appBucket = "bboy-jam-prod.appspot.com"
	sessionSecretFile = "session_secret_key"
	userSessionCookieName = "user_session"

	// Keys to store values in session.
	userIdKey = "userId"
	usernameKey = "username"
)

// TODO: Think about 2 users on the same browser.
var (
	// CookieStore saves sessions data in encrypted cookies to be stored on clients.
	// Server need not persist sessions, but can decrypt sessions data instead.
	store *sessions.CookieStore
)

func InitStore(r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "cannot create cloud storage client")
	}

	bkt := client.Bucket(appBucket)
	obj := bkt.Object(sessionSecretFile)
	reader, err := obj.NewReader(ctx)
	defer reader.Close()
	if err != nil {
		log.Errorf(ctx, "cannot obtain reader: %v", err)
		return
	}

	buf := make([]byte, reader.Attrs.Size)
	if _, err = reader.Read(buf); err != nil {
		log.Errorf(ctx, "cannot read file: %v", err)
		return
	}

	store = sessions.NewCookieStore(buf)
}

// NewUserSession creates and returns a new session representing a logged-in user.
func NewUserSession(u *sixstep.User) *Session {
	s := New(userSessionCookieName)
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
	return Get(userSessionCookieName, r)
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
	userId := s.sess.Values[userIdKey]
	if userId == nil {
		return 0, fmt.Errorf("session does not contain userId")
	}
	res, ok := userId.(int64)
	if !ok {
		return 0, fmt.Errorf("corrupt session (userId is not an int)")
	}
	return res, nil
}

// UserId returns the username stored in the session.
func (s *Session) Username() (string, error) {
	username := s.sess.Values[usernameKey]
	if username == nil {
		return "", fmt.Errorf("session does not contain username")
	}
	res, ok := username.(string)
	if !ok {
		return "", fmt.Errorf("corrupt session (username is not a string)")
	}
	return res, nil
}

