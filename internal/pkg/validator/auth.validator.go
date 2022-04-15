package validator

// Struct that define the validator/binding of Login Request
type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

// Struct that define the validator/binding of Register Request
type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	// It's kinda required, but not will included while binding
	// Avatar string
}
