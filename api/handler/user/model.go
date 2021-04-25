package user

type requestChangePassword struct {
	UserID          string `json:"user_id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type responseChangePassword struct {
}

type requestRecoveryPassword struct {
	Email    string `json:"email"`
	UserCode int64	`json:"user_code"`
}

type requestValidAndChangeCode struct {
	Email    string `json:"email"`
}

type responseRecoveryPassword struct {
}

type RequestExistEmail struct {
	Email string `json:"email"`
}
