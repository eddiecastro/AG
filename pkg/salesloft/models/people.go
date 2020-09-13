package models




type People struct {
	DisplayName      string `json:"display_name"`
	EmailAddress     string `json:"email_address"`
	ID               int64  `json:"id"`
	Title            string   `json:"title"`
}