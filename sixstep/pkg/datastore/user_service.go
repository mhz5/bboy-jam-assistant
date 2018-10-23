package datastore

import (
	"context"
	"fmt"

	"bboy-jam-assistant/sixstep/cmd/sixstep"

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

// TODO: Should you pass ctx in or not?
// User returns the User with the provided username, or an error if the lookup fails.
func (s *UserService) User(ctx context.Context, username string) (*sixstep.User, error) {
	query := datastore.NewQuery(userKind).Filter("Username = ", username)
	results := query.Run(ctx)

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

// CreateUser creates a user with the provided username and passwordHash, and saves it to datastore.e
func (s *UserService) CreateUser(ctx context.Context, username, passwordHash string) (*sixstep.User, error) {
	id, _, err := datastore.AllocateIDs(ctx, userKind, nil, 1)
	if err != nil {
		return nil, err
	}

	u := &sixstep.User{
		Id: sixstep.UserId(id),
		Username:     username,
		PasswordHash: passwordHash,
	}
	key := datastore.NewKey(ctx, userKind, "", id, nil)
	_, err = datastore.Put(ctx, key, u)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (s *UserService) DeleteUser(ctx context.Context, id sixstep.UserId) error {
	return fmt.Errorf("not implemented")
}
