package http

import (
	"testing"
)

func Test_GetUserList(t *testing.T) {
	httpRepo := New("https://run.mocky.io/v3/87931203-8086-43ef-ba16-4c8903d8fa88")
	_, err := httpRepo.GetUserList()
	if err != nil {
		t.Fail()
	}
}

func Test_GetUserList_IncorrectURL(t *testing.T) {
	httpRepo := New("https://run.mocky.io/v3/87931203-8086-43ef-ba16-4c8903d8fa89")
	_, err := httpRepo.GetUserList()
	if err == nil {
		t.Fail()
	}
	var e error = *err
	if e.Error() != "Invalid API Url" {
		t.Fail()
	}
}

func Test_GetUserList_APINotAvailableURL(t *testing.T) {
	httpRepo := New("https://run.mocky.io/v3/03d2a7bd-f12f-4275-9e9a-84e41f9c2aae")
	_, err := httpRepo.GetUserList()
	if err == nil {
		t.Fail()
	}
	var e error = *err
	if e.Error() != "Api is not available now. please try again later" {
		t.Fail()
	}
}
