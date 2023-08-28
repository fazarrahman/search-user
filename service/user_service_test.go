package service

import (
	"testing"

	"github.com/fazarrahman/search-user/domain/user/entity"
)

func Test_FilterUser(t *testing.T) {
	salaries := []entity.SalaryCSV{}
	salaries = append(salaries, entity.SalaryCSV{
		Name: "Payne",
		Tags: "euismod,proident",
	})
	salaries = append(salaries, entity.SalaryCSV{
		Name: "Max",
		Tags: "euismod,dolor",
	})
	tags := []string{"euismod", "proident"}
	res := FilterUser(salaries, tags)
	if len(res) == 0 || len(res) > 1 {
		t.Fail()
	}
}

func Test_FilterUser_NoData(t *testing.T) {
	salaries := []entity.SalaryCSV{}
	salaries = append(salaries, entity.SalaryCSV{
		Name: "Payne",
		Tags: "euismod,proident",
	})
	salaries = append(salaries, entity.SalaryCSV{
		Name: "Max",
		Tags: "euismod,dolor",
	})
	tags := []string{"euismod", "elit"}
	res := FilterUser(salaries, tags)
	if len(res) != 0 {
		t.Fail()
	}
}
