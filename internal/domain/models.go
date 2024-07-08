package domain

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUser struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,password"`
}

type UpdateUserGeneralInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,password"`
}

type ChangeEmailRequest struct {
	NewEmail string `json:"new_email" binding:"required,email"`
}

func (u *UpdateUserGeneralInfo) IsValid() bool {
	if u.FirstName == "" && u.LastName == "" && u.Nickname == "" {
		return false
	}
	return true
}

type AuthenticateUser struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}
