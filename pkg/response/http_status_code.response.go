package response

const (
	// Success Codes
	ErrCodeSuccess = 20001
	ErrCodeCreated = 20002

	// Client Error Codes
	ErrCodeBadRequest          = 40001
	ErrCodeUnauthorized        = 40002
	ErrCodeForbidden           = 40003
	ErrCodeNotFound            = 40004
	ErrCodeConflict            = 40005 // Mã lỗi khi xảy ra xung đột (conflict)
	ErrCodeUnprocessableEntity = 40006 // Dữ liệu không thể xử lý

	// Server Error Codes
	ErrCodeInternalError      = 50001
	ErrCodeServiceUnavailable = 50002
	ErrCodeGatewayTimeout     = 50003 // Thời gian chờ quá lâu

	// Authentication Errors
	ErrInvalidToken       = 40101
	ErrTokenExpired       = 40102
	ErrTokenRevoked       = 40103
	ErrInvalidCredentials = 40104
	ErrUserExists         = 40105 // Người dùng đã tồn tại

	// Database Errors
	ErrDBConnectionFail = 50006
	ErrDBQueryFail      = 50007
	ErrDBInsertFail     = 50008
	ErrDBUpdateFail     = 50009
	ErrDBDeleteFail     = 50010
)

var msg = map[int]string{
	// Success
	ErrCodeSuccess: "Success",
	ErrCodeCreated: "Resource created successfully",

	// Client Errors
	ErrCodeBadRequest:          "Bad request: The server could not understand the request due to invalid syntax.",
	ErrCodeUnauthorized:        "Unauthorized: Authentication is required and has failed or has not yet been provided.",
	ErrCodeForbidden:           "Forbidden: The request was valid, but the server is refusing action.",
	ErrCodeNotFound:            "Not Found: The requested resource could not be found.",
	ErrCodeConflict:            "Conflict: The request could not be completed due to a conflict with the current state of the resource.",
	ErrCodeUnprocessableEntity: "Unprocessable Entity: The request was well-formed but was unable to be followed due to semantic errors.",

	// Server Errors
	ErrCodeInternalError:      "Internal Server Error: An unexpected condition was encountered.",
	ErrCodeServiceUnavailable: "Service Unavailable: The server is currently unable to handle the request due to temporary overloading or maintenance.",
	ErrCodeGatewayTimeout:     "Gateway Timeout: The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.",

	// Authentication Errors
	ErrInvalidToken:       "Invalid Token: The provided token is invalid.",
	ErrTokenExpired:       "Token Expired: The provided token has expired.",
	ErrTokenRevoked:       "Token Revoked: The token has been revoked and is no longer valid.",
	ErrInvalidCredentials: "Invalid Credentials: The username or password is incorrect.",
	ErrUserExists:         "User Already Exists: A user with the same credentials already exists.",

	// Database Errors
	ErrDBConnectionFail: "Database Connection Failed: Unable to connect to the database.",
	ErrDBQueryFail:      "Database Query Failed: Unable to execute the query.",
	ErrDBInsertFail:     "Database Insert Failed: Unable to insert the record.",
	ErrDBUpdateFail:     "Database Update Failed: Unable to update the record.",
	ErrDBDeleteFail:     "Database Delete Failed: Unable to delete the record.",
}
