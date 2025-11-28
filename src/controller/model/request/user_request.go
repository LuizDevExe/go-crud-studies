package request

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%^&*"`
	Age      int8   `json:"age" binding:"required,gte=10,lte=130"`
}
