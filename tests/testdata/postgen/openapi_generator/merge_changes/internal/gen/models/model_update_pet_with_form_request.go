// Code generated by openapi-generator. DO NOT EDIT.

package models

type UpdatePetWithFormRequest struct {
	// Updated name of the pet.
	Name string `json:"name,omitempty"`
	// Updated status of the pet.
	Status string `json:"status,omitempty"`
}

// TODO validate everything, accumulate errors and return error map instead.
// validate ...
func (o *UpdatePetWithFormRequest) validate() error {
	return nil
}
