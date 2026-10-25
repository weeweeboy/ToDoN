package users_postgres_repository

import "ToDoN/internal/core/domain"

type UserModel struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func UserDomainsFromModels(models []UserModel) []domain.User {
	userDomains := make([]domain.User, len(models))
	for i, user := range models {
		userDomains[i] = domain.NewUser(user.ID, user.Version, user.FullName, user.PhoneNumber)
	}
	return userDomains
}

func UserDomainFromModel(models UserModel) domain.User {
	userDomain := domain.NewUser(models.ID, models.Version, models.FullName, models.PhoneNumber)
	return userDomain
}
