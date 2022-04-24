package dtos

type RegisterDTO struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,len=10"`
	Password  string `json:"password" validate:"required,gte=8"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type ForgetPasswordDTO struct {
	Email       string `json:"email" validate:"required,email"`
	NewPassword string `json:"new_password" validate:"required,gte=8"`
}

type CheckEmailVerificationDTO struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code"`
}
