package http

// CreateRequest represents the HTTP payload for creating an organisation.
type CreateRequest struct {
	Name string  `json:"name"`
	Desc *string `json:"description,omitempty"`
}

// UpdateRequest represents the HTTP payload for updating an organisation.
type UpdateRequest struct {
	Name *string `json:"name,omitempty"`
	Desc *string `json:"description,omitempty"`
}
