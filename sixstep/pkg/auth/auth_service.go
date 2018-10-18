package auth

import (
	"context"
	"fmt"

	"bboy-jam-assistant/sixstep/cmd/sixstep"
	"bboy-jam-assistant/sixstep/pkg/datastore"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	// TODO: Does auth service need a reference to user service?
	userService sixstep.UserService
}

func NewService() sixstep.AuthService {
	return &Service{
		userService: datastore.NewUserService(),
	}
}

func (s *Service) Authenticate(ctx context.Context, username string, password string) (*sixstep.User, error) {
	u, err := s.userService.User(ctx, username)
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	return u, nil
}

func (s *Service) Authorize() {}

