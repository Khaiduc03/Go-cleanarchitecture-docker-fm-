package modelFeedback

type GetFeedbackRes struct {
	ID             uint     `json:"id"`
	Name_Feed_Back string   `json:"name_feed_back" validate:"required"`
	RoomID         uint     `json:"room_id" validate:"required,gte=1"`
	Description    string   `json:"description" validate:"required"`
	CategoryID     uint     `json:"category_id" validate:"required,gte=1"`
	UserID         uint     `json:"user_id" validate:"required,gte=1"`
	Urls           []string `json:"url" validate:"required"`
}

type GetAllFeedbackRes struct {
	ID             uint   `json:"id"`
	Name_Feed_Back string `json:"name_feed_back" validate:"required"`
	RoomID         uint   `json:"room_id" validate:"required,gte=1"`
	Description    string `json:"description" validate:"required"`
	CategoryID     uint   `json:"category_id" validate:"required,gte=1"`
	UserID         uint   `json:"user_id" validate:"required,gte=1"`
}
