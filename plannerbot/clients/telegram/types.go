package telegram

type Update struct {
	updateID int    `json:"update_id"`
	message  string `json:"message"`
}

type UpdatesResponce struct {
	OK     bool     `json:"ok"`
	result []Update `json:"result"`
}
