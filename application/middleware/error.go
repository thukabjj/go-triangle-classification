package middleware

type Error struct {
	Code    int               `json:"code"`
	Status  string            `json:"status"`
	Details map[string]string `json:"datails"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}
