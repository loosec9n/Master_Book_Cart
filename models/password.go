package models

type ForgotPasswordInput struct {
	Email string `json:"forgetEmail"`
}

type ResetPasswordInput struct {
	Email           string `json:"email,omitempty"`
	TokenPassword   string `json:"tokenPassword,omitempty"`
	PasswordConfirm string `json:"resetPasswordConfirm,omitempty"`
}
