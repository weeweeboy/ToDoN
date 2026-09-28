package users_service

import (
	"ToDoN/internal/core/domain"
	"context"
	"fmt"
)

func (s *UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	user, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
