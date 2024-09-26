package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
)

var mgs = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Invalid",
}
