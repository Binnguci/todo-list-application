package error

import "fmt"

type AppError struct {
	Code    int
	Message string
}

var (
	SUCCESS        = AppError{200, "Success"}
	DELETE_SUCCESS = AppError{200, "Delete success"}
	CREATE_SUCCESS = AppError{201, "Create success"}
	UPDATE_SUCCESS = AppError{200, "Update success"}
	FOUND          = AppError{200, "Found"}

	INVALID_REQUEST         = AppError{400, "Invalid request"}
	PARAMETER_NOT_VALID     = AppError{400, "Parameter not valid or validation error"}
	PARAMETER_MISSING       = AppError{400, "Parameter missing"}
	DELETE_FAILED           = AppError{400, "Delete failed"}
	CREATE_FAILED           = AppError{400, "Create failed"}
	UPDATE_FAILED           = AppError{400, "Update failed"}
	UNAUTHENTICATED         = AppError{401, "Unauthenticated"}
	UNAUTHORIZED            = AppError{403, "You do not have permission"}
	USER_NOT_FOUND          = AppError{404, "User not found"}
	NOT_FOUND               = AppError{404, "Not found"}
	EMAIL_ALREADY_EXISTS    = AppError{409, "Email already exists"}
	USERNAME_ALREADY_EXISTS = AppError{409, "Username already exists"}
	INVALID_PASSWORD        = AppError{400, "Password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character"}
	OTP_EXPIRED             = AppError{400, "OTP expired"}
	INVALID_OTP             = AppError{400, "Invalid OTP"}
	ACCOUNT_NOT_VERIFIED    = AppError{400, "Account not verified"}

	INTERNAL_SERVER_ERROR = AppError{500, "Internal Server Error"}
	SERVICE_UNAVAILABLE   = AppError{503, "Service Unavailable"}
	GATEWAY_TIMEOUT       = AppError{504, "Gateway Timeout"}
)

func (e AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
