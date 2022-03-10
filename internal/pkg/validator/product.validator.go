package validator

// Struct that define the validator/binding of Create Product Request
type CreateProductRequest struct {
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Price uint64 `json:"price" form:"email" binding:"required"`
}

// Struct that define the validator/binding of Update Product Request
type UpdateProductRequest struct {
	ID    int64  `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Price uint64 `json:"price" form:"email" binding:"required"`
}
