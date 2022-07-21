package models

type ForgotPasswordInput struct {
	Email string `json:"forgetEmail"`
}

type ResetPasswordInput struct {
	Email           string `json:"email"`
	TokenPassword   string `json:"tokenPassword"`
	PasswordConfirm string `json:"resetPasswordConfirm"`
}
