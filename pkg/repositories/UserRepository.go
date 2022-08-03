package repositories

import (
	"github.com/go-crud/pkg/forms"
	"github.com/go-crud/pkg/models"
)

type UserRepository struct {
	Repository DefaultRepository
}

func (r UserRepository) Index() (users []models.User, err error) {
	db := r.Repository.db()

	err = db.Find(&users).Error
	return
}

func (r UserRepository) Show(id string) (user models.User, err error) {
	db := r.Repository.db()

	err = db.First(&user, id).Error
	return
}

func (r UserRepository) Store(form forms.StoreUserForm) (user models.User, err error) {
	db := r.Repository.db()

	user = models.User{}
	user.Name = form.Name
	user.Email = form.Email
	user.Password, err = user.HashPassword(form.Password)

	db.Save(&user)

	return
}

func (r UserRepository) Update(id string, form forms.UpdateUserForm) (user models.User, err error) {
	db := r.Repository.db()

	err = db.Find(&user, id).Error

	if err != nil {
		return
	}

	user.Name = form.Name
	user.Email = form.Email
	user.Password, err = user.HashPassword(form.Password)

	db.Save(&user)

	return
}

func (r UserRepository) Destroy(id string) (err error) {
	db := r.Repository.db()

	err = db.Delete(&models.User{}, id).Error

	return
}
