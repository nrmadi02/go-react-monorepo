package dto

import (
	"regexp"
	"time"
)

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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
