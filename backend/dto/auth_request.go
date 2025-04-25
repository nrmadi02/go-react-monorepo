package dto

import "regexp"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r LoginRequest) ValidationMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"Email": {
			"required": "Email is required",
			"email":    "Email is not valid",
		},
		"Password": {
			"required": "Password is required",
		},
	}
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (r RegisterRequest) Validate() []string {
	errs := []string{}
	if matched, _ := regexp.MatchString(`[^a-zA-Z0-9 ]`, r.Name); matched {
		errs = append(errs, "Name not allowed special characters")
	}
	if regexp.MustCompile(` {2,}`).MatchString(r.Name) {
		errs = append(errs, "Name not allowed more than two spaces")
	}
	return errs
}

func (r RegisterRequest) ValidationMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"Name": {
			"required": "Name is required",
		},
		"Email": {
			"required": "Email is required",
			"email":    "Email is not valid",
		},
		"Password": {
			"required": "Password is required",
		},
	}
}
