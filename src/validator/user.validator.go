package validator

type UserUpdateValidator struct {
	Name     string `json:"name,omitempty" form:"name,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" form:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required,min:6"`
}

type UserCreateValidator struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min:6"`
}
