package twitter

// Tweet ...
type Tweet struct {
	CreatedAt     string `json:"created_at"`
	FavoriteCount int    `json:"favorite_count"`
	Favorited     bool   `json:"favorited"`
	FilterLevel   string `json:"filter_level"`
	ID            int64  `json:"id"`
	IDSting       string `json:"id_str"`
	InReplyToName string `json:"in_reply_to_screen_name"`
	InReplyToID   int64  `json:"in_reply_to_status_id"`
	InReplyToUser string `json:"in_reply_to_user_id_str"`
	Text          string `json:"text"`
}
