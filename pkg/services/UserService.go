package services

import (
	"errors"
	"github.com/go-crud/pkg/forms"
	"github.com/go-crud/pkg/models"
	"github.com/go-crud/pkg/repositories"
)

type UserService struct{}

var r repositories.UserRepository

func init() {
	r = repositories.UserRepository{}
}

func (s UserService) Index() (users []models.User, err error) {
	users, err = r.Index()

	return
}

func (s UserService) Show(id string) (user models.User, err error) {
	user, err = r.Show(id)

	return
}

func (s UserService) Store(form forms.StoreUserForm) (user models.User, err error) {
	form.Password, err = user.HashPassword(form.Password)

	user, err = r.Store(form)

	return
}

func (s UserService) Update(id string, form forms.UpdateUserForm) (user models.User, err error) {
	user, err = r.Update(id, form)

	return
}

func (s UserService) UpdatePassword(id string, form forms.UpdateUserPassword) (err error) {
	if form.Password != form.ConfirmPassword {
		err = errors.New("Passwords do not match")
		return
	}

	form.Password, err = models.User{}.HashPassword(form.Password)

	err = r.UpdatePassword(id, form)

	return
}

func (s UserService) Delete(id string) (err error) {
	err = r.Destroy(id)

	return
}
