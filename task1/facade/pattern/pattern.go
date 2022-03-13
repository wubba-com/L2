package facade

import (
	"github.com/go-playground/validator/v10"
)

func NewValidaterFacade() *validaterFacade {
	v := validator.New()
	return &validaterFacade{validater: v}
}

type validaterFacade struct {
	validater *validator.Validate
}

func (v *validaterFacade) ValidateStruct(val interface{}) error {
	err := v.validater.Struct(val)
	if err != nil {
		return err

	}
	return nil
}

func (v *validaterFacade) IsEmail(email string) error {
	err := v.validater.Var(email, "required,email")
	if err != nil {
		return err
	}
	return nil
}
