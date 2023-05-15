package dto

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=3"`
}

func (s SignInInput) Validate() error {
	return validate.Struct(s)
}
