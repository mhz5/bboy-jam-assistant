package datastore

import (
	"bboy-jam-assistant/sixstep/cmd/sixstep"
	"context"
	"fmt"

	"google.golang.org/appengine/datastore"
)

const (
	userKind = "User"
)

// TODO: Can we make aliases for the domain types?
var (
	User   sixstep.User
	UserId sixstep.UserId
)

type UserService struct {
}

var _ sixstep.UserService = &UserService{}

func NewUserService() *UserService {
	return &UserService{}
}

// TODO: Should you pass ctx in?
// TODO: Change signature to (ctx, username, passwordHash)?
func (s *UserService) User(ctx context.Context, username string) (*sixstep.User, error) {
	query := datastore.NewQuery(userKind).Filter("Username = ", username)
	results := query.Run(ctx)

	// TODO: Fix issue where user doesn't have an ID.
	user := &sixstep.User{}
	_, err := results.Next(user)
	if err == datastore.Done {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, username, password string) (*sixstep.User, error) {
	u := &sixstep.User{
		Username:     username,
		PasswordHash: password,
	}
	key := datastore.NewIncompleteKey(ctx, userKind, nil)
	key, err := datastore.Put(ctx, key, u)
	u.Id = sixstep.UserId(key.IntID())

	if err != nil {
		return nil, err
	}
	return u, err
}

func (s *UserService) DeleteUser(ctx context.Context, id sixstep.UserId) error {
	return nil
}
