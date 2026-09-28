package users_service

import (
	"ToDoN/internal/core/domain"
	core_errors "ToDoN/internal/core/errors"
	"context"
	"fmt"
)

func (u *UserService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be > 0 : %w", core_errors.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be > 0 : %w", core_errors.ErrInvalidArgument)
	}

	users, err := u.userRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users from repo : %w", err)
	}
	return users, nil
}
