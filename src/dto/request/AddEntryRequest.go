package request

type AddEntryRequest struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}
