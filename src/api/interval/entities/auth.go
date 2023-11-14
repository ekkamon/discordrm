package entities

type LoginWithPasswordRequest struct {
	Username string `json:"username" validate:"required,ascii,gte=6,lte=255"`
	Password string `json:"password" validate:"required,ascii,gte=6,lte=255"`
}

type TokenData struct {
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}
