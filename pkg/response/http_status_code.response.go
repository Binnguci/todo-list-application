package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
)

var mgs = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Invalid",
	ErrInvalidToken:     "Token invalid",
}
