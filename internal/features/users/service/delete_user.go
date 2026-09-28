package users_service

import (
	"context"
	"fmt"
)

func (u *UserService) DeleteUser(ctx context.Context, userID int) error {
	err := u.userRepository.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
