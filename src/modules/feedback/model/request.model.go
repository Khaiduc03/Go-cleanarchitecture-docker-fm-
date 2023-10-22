package modelFeedback



type CreateFeedbackReq struct {
	NameFeedBack string `json:"name_feed_back"`
	RoomID       int    `json:"room_id"`
	UserID       int    `json:"user_id"`
	Description  string `json:"description"`
	CategoryID   []int  `json:"category_id"`
	
}