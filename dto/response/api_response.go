package response

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Thêm "omitempty" để ẩn khi Data là null
	Error   string      `json:"error,omitempty"`
}
