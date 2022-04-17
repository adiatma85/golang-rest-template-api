package validator

// Struct that define the validator/binding of Register Request (Admin)
type RegisterNewUserRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	UserType string `json:"user_type" form:"user_type"`
}

// Struct that define the validator/binding of Update User Request
type UpdateUserRequest struct {
	Name     string `json:"name" form:"name" validation:"name"`
	Email    string `json:"email" form:"email" validation:"email"`
	Password string `json:"password" form:"password" validation:"password"`
}

// Struct that define the validator/binding of Delete Users by JSON Request
type DeleteUsersRequest struct {
	Ids []uint64 `json:"ids" form:"ids" binding:"required"`
}
