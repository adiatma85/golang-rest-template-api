package validator

// Struct that define the validator/binding of Update User Request
type UpdateUserRequest struct {
	ID    int64  `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Email string `json:"email" form:"email" binding:"required,email"`
}
