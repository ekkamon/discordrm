package langs

// common errors
const (
	ErrNotFound       = "error_not_found"
	ErrUnAuthorized   = "error_unauthorized"
	ErrBadRequest     = "error_bad_request"
	ErrInternalServer = "error_internal_server"
)

// user errors
const (
	ErrUsernameAlreadyExists = "error_username_already_exists"
	ErrEmailAlreadyExists    = "error_email_already_exists"
)
