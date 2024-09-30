package response

const (
	// Success
	ErrCodeSuccess = 20001
	ErrCodeCreated = 20002

	// Request Errors
	ErrCodeParamInvalid = 20003
	ErrCodeNotFound     = 20004
	ErrCodeUnauthorized = 20005
	ErrCodeForbidden    = 20006
	ErrCodeBadRequest   = 20007

	// Authentication/Token Errors
	ErrInvalidToken       = 30001
	ErrTokenExpired       = 30002
	ErrTokenRevoked       = 30003
	ErrInvalidCredentials = 30004

	// Database Errors
	ErrDBConnectionFail = 40001
	ErrDBQueryFail      = 40002
	ErrDBInsertFail     = 40003
	ErrDBUpdateFail     = 40004
	ErrDBDeleteFail     = 40005

	// System Errors
	ErrSystemInternal     = 50001
	ErrServiceUnavailable = 50002
	ErrTimeout            = 50003
)

var msg = map[int]string{
	// Success
	ErrCodeSuccess: "Success",
	ErrCodeCreated: "Resource created successfully",

	// Request Errors
	ErrCodeParamInvalid: "Invalid parameters",
	ErrCodeNotFound:     "Resource not found",
	ErrCodeUnauthorized: "Unauthorized access",
	ErrCodeForbidden:    "Access forbidden",
	ErrCodeBadRequest:   "Bad request",

	// Authentication/Token Errors
	ErrInvalidToken:       "Invalid token",
	ErrTokenExpired:       "Token expired",
	ErrTokenRevoked:       "Token has been revoked",
	ErrInvalidCredentials: "Invalid credentials",

	// Database Errors
	ErrDBConnectionFail: "Database connection failed",
	ErrDBQueryFail:      "Failed to execute query",
	ErrDBInsertFail:     "Failed to insert record",
	ErrDBUpdateFail:     "Failed to update record",
	ErrDBDeleteFail:     "Failed to delete record",

	// System Errors
	ErrSystemInternal:     "Internal server error",
	ErrServiceUnavailable: "Service unavailable",
	ErrTimeout:            "Request timed out",
}
