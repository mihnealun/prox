package response

// Base response skeleton
type Base struct {
	Code   int      `json:"-"`
	Errors []string `json:"errors"`
}
