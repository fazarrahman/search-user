package service

import (
	"log"
)

func (s *Svc) GetUserList() {
	salaries, err := s.userRepo.GetUserList()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(salaries)
}
