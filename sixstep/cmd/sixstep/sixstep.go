package sixstep

import (
	"context"
)
// Domain types

// TODO: How to make aliases without needing to cast type in every subpackage?
type UserId int64

type User struct {
	Id           UserId
	Username     string
	PasswordHash string
}

type UserService interface {
	User(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, id UserId) error
}

type AuthService interface {
	Authenticate(ctx context.Context, username string, password string) (*User, error)
	// TODO
	Authorize()
}

type Router interface {
	// TODO
	Handle()
}

type Server interface {
	// TODO
	Serve()
}
