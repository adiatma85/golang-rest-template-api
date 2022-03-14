package validator

// Struct that define the validator/binding of Update User Request
type UpdateUserRequest struct {
	Name     string `json:"name" form:"name" validation:"name"`
	Email    string `json:"email" form:"email" validation:"email"`
	Password string `json:"password" form:"password" validation:"password"`
}
