package models

type Update struct {
	Ok     bool       `json:"ok"`
	Result ResultList `json:"result"`
}
type ResultList []struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}
type Result struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}
type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}

type Message struct {
	MessageID int    `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}
type From struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}
type DtoRefQuery struct {
	ID              int64  `json:"id"`
	Quries          string `json:"quries"`
	ExpectedMessage string `json:"expected_message"`
}
