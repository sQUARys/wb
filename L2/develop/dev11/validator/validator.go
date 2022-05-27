package validator

import (
	"dev11/structs"
	"fmt"
)

//Validation
type Validator struct {
	err error
}

func IsValid(option string, i structs.Input) error {
	validator := Validator{}

	switch option {
	case "del":
		validator.MustHaveDate(i)
	case "crt":
		validator.MustHaveId(i)
		validator.MustHaveName(i)
	case "upd":
		validator.MustHaveNameOrId(i)
		validator.MustHaveDate(i)
	}
	return validator.IsValid()
}

func (v *Validator) IsValid() error {
	return v.err
}

func (v *Validator) MustHaveId(input structs.Input) bool {
	if v.err != nil {
		return false
	}
	if input.ID == "" {
		v.err = fmt.Errorf("Must have an id field")
		return false
	}
	return true
}

func (v *Validator) MustHaveName(input structs.Input) bool {
	if v.err != nil {
		return false
	}
	if input.Name == "" {
		v.err = fmt.Errorf("Must have a name field")
		return false
	}
	return true
}

func (v *Validator) MustHaveNameOrId(input structs.Input) bool {
	if v.err != nil {
		return false
	}
	if input.Name == "" && input.ID == "" {
		v.err = fmt.Errorf("Must have a name or id fields")
		return false
	}
	return true
}

func (v *Validator) MustHaveDate(input structs.Input) bool {
	if v.err != nil {
		return false
	}
	if input.Date == "" {
		v.err = fmt.Errorf("Must have a date field")
		return false
	}
	return true
}
