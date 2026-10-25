package users_postgres_repository

import (
	"ToDoN/internal/core/domain"
	core_errors "ToDoN/internal/core/errors"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (u *UsersRepository) PatchUser(ctx context.Context, userID int, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.pool.OpTimeout())
	defer cancel()

	query := `UPDATE todoapp.users
          SET full_name=$1,
              phone_number=$2,
              version=version+1
          WHERE id=$3 AND version=$4
          RETURNING id, version, full_name, phone_number`

	row := u.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, userID, user.Version)

	var userModel UserModel

	err := row.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id '%d' concurrently accessed:  %w", userID, core_errors.ErrConflict)
		}
		return domain.User{}, fmt.Errorf("failed to scan patched user: %w", err)

	}
	response := UserDomainFromModel(userModel)
	return response, nil
}
