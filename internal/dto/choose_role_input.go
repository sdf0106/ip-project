package dto

type ChangeRoleInput struct {
	PrevRole string `json:"prev_role" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

func (c ChangeRoleInput) Validate() error {
	return validate.Struct(c)
}
