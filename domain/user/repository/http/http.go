package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/fazarrahman/search-user/domain/user/entity"
)

type UserAPIHttp struct {
	Url string
}

func New(url string) *UserAPIHttp {
	return &UserAPIHttp{Url: url}
}

func (u *UserAPIHttp) GetUserList() ([]*entity.Salary, *error) {
	res, err := http.Get(u.Url)
	if err != nil {
		return nil, &err
	}
	if res.StatusCode >= 500 {
		errr := errors.New("Api is not available now. please try again later")
		return nil, &errr
	} else if res.StatusCode == 404 {
		errr := errors.New("Invalid API Url")
		return nil, &errr
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &err
	}

	var salaries []*entity.Salary
	err = json.Unmarshal(body, &salaries)
	if err != nil {
		return nil, &err
	}

	return salaries, nil
}
