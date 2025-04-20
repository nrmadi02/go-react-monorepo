package dto

import "regexp"

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// Validate implements custom validation for CreateUserRequest
func (r CreateUserRequest) Validate() []string {
	errs := []string{}
	if matched, _ := regexp.MatchString(`[^a-zA-Z0-9 ]`, r.Name); matched {
		errs = append(errs, "Name not allowed special characters")
	}
	if regexp.MustCompile(` {2,}`).MatchString(r.Name) {
		errs = append(errs, "Name not allowed more than two spaces")
	}
	return errs
}

func (r CreateUserRequest) ValidationMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"Name": {
			"required": "Name is required",
		},
		"Email": {
			"required": "Email is required",
			"email":    "Email is not valid",
		},
	}
}
