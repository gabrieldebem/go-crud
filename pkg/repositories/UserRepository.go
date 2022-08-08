package repositories

import (
	"github.com/go-crud/pkg/forms"
	"github.com/go-crud/pkg/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository DefaultRepository
}

var db *gorm.DB

func init() {
	db = UserRepository{}.Repository.db()
}

func (r UserRepository) Index() (users []models.User, err error) {
	err = db.Find(&users).Error
	return
}

func (r UserRepository) Show(id string) (user models.User, err error) {
	err = db.First(&user, id).Error
	return
}

func (r UserRepository) Store(form forms.StoreUserForm) (user models.User, err error) {
	user = models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	err = db.Create(&user).Error

	return
}

func (r UserRepository) Update(id string, form forms.UpdateUserForm) (user models.User, err error) {
	err = db.Model(&user).Where("id = ?", id).Updates(models.User{
		Name:  form.Name,
		Email: form.Email,
	}).Error

	return
}

func (r UserRepository) UpdatePassword(id string, form forms.UpdateUserPassword) (err error) {
	err = db.Model(&models.User{}).Where(
		"id = ?",
		id,
	).Update(
		"password",
		form.Password,
	).Error

	return
}

func (r UserRepository) Destroy(id string) (err error) {
	err = db.Delete(&models.User{}, id).Error

	return
}
