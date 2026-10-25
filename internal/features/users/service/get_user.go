package users_service

import (
	"ToDoN/internal/core/domain"
	"context"
	"fmt"
)

func (u *UserService) GetUser(ctx context.Context, userID int) (domain.User, error) {
	user, err := u.userRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user from repo: %w", err)
	}

	return user, nil
}
