package t_bot

type User struct {
	ID         int    `json:"id"`
	ExternalID string `json:"external_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	UserName   string `json:"user_name"`
	Score      int64  `json:"score"`
}

type UserFilter struct {
	Limit int
}
