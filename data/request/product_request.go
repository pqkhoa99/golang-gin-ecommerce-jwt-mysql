package request

type CreateProductRequest struct {
	Name  string  `validate:"required,min=2,max=100" json:"name"`
	Price float64 `validate:"required"`
	Stock int     `validate:"required"`
}

type UpdateProductRequest struct {
	Id    int     `validate:"required"`
	Name  string  `validate:"required,min=2,max=100" json:"username"`
	Price float64 `validate:"required"`
	Stock int     `validate:"required"`
}
