package users_postgres_repository

import (
	"ToDoN/internal/core/domain"
	core_errors "ToDoN/internal/core/errors"
	core_postgres_pool "ToDoN/internal/core/repository/postgres/pool"
	"context"
	"errors"
	"fmt"
)

func (u *UsersRepository) GetUser(ctx context.Context, userID int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.pool.OpTimeout())
	defer cancel()
	query := `SELECT id, version, full_name, phone_number FROM todoapp.users WHERE id = $1`
	row := u.pool.QueryRow(ctx, query, userID)

	var userModel UserModel
	err := row.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id %d not found : %w", userID, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("failed to scan user: %w", err)
	}
	user := UserDomainFromModel(userModel)
	return user, nil
}
