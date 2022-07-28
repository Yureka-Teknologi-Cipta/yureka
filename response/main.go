package response

type Base struct {
	Status     int               `json:"status"`
	Message    string            `json:"message"`
	Validation map[string]string `json:"validation"`
	Data       interface{}       `json:"data"`
}

type List struct {
	List  []interface{} `json:"list"`
	Limit int64         `json:"limit"`
	Page  int64         `json:"page"`
	Total int64         `json:"total"`
}
