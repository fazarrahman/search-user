package service

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fazarrahman/search-user/domain/user/entity"
	"github.com/gocarina/gocsv"
)

func (s *Svc) SearchUser(tags []string) {
	// Open the CSV file
	file, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Please fetch the user list first")
	}
	defer file.Close()

	// Read the CSV file into a slice of Salary structs
	var salaryCSV []entity.SalaryCSV
	if err := gocsv.UnmarshalFile(file, &salaryCSV); err != nil {
		log.Fatalln(err)
	}

	// Print the Users
	res := FilterUser(salaryCSV, tags)
	for _, r := range res {
		fmt.Printf("Name: %s, Salary: %s\n", r.Name, r.Balance)
	}
}

func FilterUser(salaries []entity.SalaryCSV, tags []string) []entity.SalaryCSV {
	var isUserHaveTags bool = true
	var res []entity.SalaryCSV
	for _, s := range salaries {
		isUserHaveTags = true
		tagArr := strings.Split(s.Tags, ",")
		for _, t := range tags {
			if !ArrayContains(tagArr, t) {
				isUserHaveTags = false
				break
			}
		}
		if isUserHaveTags {
			res = append(res, s)
		}
	}
	return res
}

func ArrayContains(array []string, s string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == s {
			return true
		}
	}
	return false
}

func (s *Svc) GetUserList() {
	salaries, err := s.userRepo.GetUserList()
	if err != nil {
		log.Fatalln(*err)
	}

	// create csv file
	file, errl := os.Create("users.csv")
	defer file.Close()
	if errl != nil {
		log.Fatalln("failed to open file", err)
	}
	var salaryCSV []entity.SalaryCSV
	for _, s := range salaries {
		// Only write active users
		if s.IsActive {
			for _, f := range s.Friends {
				tags := strings.Join(s.Tags, ",")
				salaryCSV = append(salaryCSV, entity.SalaryCSV{
					Index:    s.Index,
					Guid:     s.Guid,
					IsActive: s.IsActive,
					Balance:  s.Balance,
					Tags:     tags,
					Name:     f.Name,
				})
			}
		}
	}

	if err := gocsv.MarshalFile(&salaryCSV, file); err != nil {
		log.Fatalln(err)
	}
}
