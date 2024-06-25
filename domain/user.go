package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	id    int
	email string
	name  string
}

func NewUser(id int, email string, name string) *User {
	return &User{
		id:    id,
		email: email,
		name:  name,
	}
}

func (i *User) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(
			&i.name,
			validation.Required.Error("名前は必須入力です"),
		),
		validation.Field(
			&i.email,
			validation.Required.Error("メールアドレスは必須入力です"),
			is.Email.Error("メールアドレスを入力して下さい"),
		),
	)
}
