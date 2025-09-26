package http

// Response represents the HTTP response model for an organisation.
type Response struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Desc *string `json:"description,omitempty"`
}
