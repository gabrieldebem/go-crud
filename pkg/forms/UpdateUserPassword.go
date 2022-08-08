package forms

type UpdateUserPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
