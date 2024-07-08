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
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,password"`
}

type UserAuthDTO struct {
	Nickname string `json:"nickname" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}
