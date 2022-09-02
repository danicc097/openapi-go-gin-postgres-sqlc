// Code generated by openapi-generator. DO NOT EDIT.

package models

// CreateUserRequest represents a new user.
type CreateUserRequest struct {
	Username string `json:"username,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// TODO validate everything, accumulate errors and return error map instead.
// validate ...
func (o *CreateUserRequest) validate() error {
	return nil
}
