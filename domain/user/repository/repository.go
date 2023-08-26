package repository

import (
	"github.com/fazarrahman/search-user/domain/user/entity"
)

type UserInterface interface {
	GetUserList() ([]*entity.Salary, *error)
}
