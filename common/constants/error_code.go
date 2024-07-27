package constants

const (
	ErrorCodeInternalServerError = "err.system.500"
)

// auth related error code
const (
	ErrorCodeUserAlreadyExists = "err.auth.10001"
	ErrorCodeUserNotFound      = "err.auth.10002"
	ErrorCodeInvalidPassword   = "err.auth.10003"
)

// user related error code
const (
	ErrorCodeUserAlreadyDeleted = "err.user.10001"
	ErrorCodeUserNotActive      = "err.user.10002"
)
