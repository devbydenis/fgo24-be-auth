package models

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email"`
	Otp      string `json:"otp"`
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
