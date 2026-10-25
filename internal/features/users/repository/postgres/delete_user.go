package users_postgres_repository

import (
	"context"
	"fmt"
)

func (u *UsersRepository) DeleteUser(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, u.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM todoapp.users WHERE id = $1`
	_, err := u.pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user repo: %w", err)
	}
	return nil
}
