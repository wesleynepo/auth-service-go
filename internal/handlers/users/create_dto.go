package users

type BodyCreateUser struct {
    Email string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
    ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
