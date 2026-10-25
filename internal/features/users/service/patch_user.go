package users_service

import (
	"ToDoN/internal/core/domain"
	"context"
	"fmt"
)

func (u *UserService) PatchUser(ctx context.Context, userID int, patch domain.UserPatch) (domain.User, error) {
	user, err := u.userRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("GetUser: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}

	patchedUser, err := u.userRepository.PatchUser(ctx, userID, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to patch user: %w", err)
	}

	return patchedUser, nil
}
