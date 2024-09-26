package request

type TaskRequest struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description"`
	UserID      uint   `json:"user_id" form:"user_id" binding:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	Deadline    string `json:"deadline" form:"deadline"`
	Tags        []uint `json:"tags" form:"tags"`
}
