package service

import userRepository "github.com/fazarrahman/search-user/domain/user/repository"

type Svc struct {
	userRepo userRepository.UserInterface
}

func New(userRepo userRepository.UserInterface) *Svc {
	return &Svc{userRepo: userRepo}
}

type Service interface {
	GetUserList()
	SearchUser(tags []string)
}
