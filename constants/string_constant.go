package constants

const (
	CONNECT_DB_SUCCESS = "Successfully connected to database"
	REGEX_PASSWORD     = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{6,20}$"
	PASSWORD_INVALID   = "Password must be between 6 and 20 characters long, include at least one uppercase letter, one lowercase letter, one digit, and one special character"
)
