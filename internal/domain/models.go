package domain

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserDTO struct {
	FirstName string `json:"first_name" validate:"required,min=3"`
	LastName  string `json:"last_name" validate:"required,min=3"`
	Nickname  string `json:"nickname" validate:"required,min=3"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=64"`
}

type UserAuthDTO struct {
	Nickname string `json:"nickname" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}
