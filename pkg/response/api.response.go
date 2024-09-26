package response

type APIResponse struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
