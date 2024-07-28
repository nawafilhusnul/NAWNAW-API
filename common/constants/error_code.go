package constants

const (
	ErrorCodeInternalServerError = "err.system.500"
	ErrorCodeBadRequest          = "err.system.400"
)

// auth related error code
const (
	ErrorCodeUserAlreadyExists = "err.auth.10001"
	ErrorCodeUserNotFound      = "err.auth.10002"
	ErrorCodeInvalidPassword   = "err.auth.10003"
	ErrorCodeMissingToken      = "err.auth.10004"
	ErrorCodeInvalidToken      = "err.auth.10005"
	ErrorCodeInvalidTimezone   = "err.auth.10006"
	ErrorCodeForbidden         = "err.auth.10007"
)

// user related error code
const (
	ErrorCodeUserAlreadyDeleted = "err.user.10001"
	ErrorCodeUserNotActive      = "err.user.10002"
)

// module related error code
const (
	ErrorCodeModuleNotFound = "err.module.10001"
)
