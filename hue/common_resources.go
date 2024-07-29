package hue

// Error represents an error with a human-readable explanation.
type Error struct {
	Description string `json:"description"`
}

// ResourceIdentifier represents a reference to another resource.
type ResourceIdentifier struct {
	RID   string `json:"rid" validate:"required"`
	RType string `json:"rtype" validate:"required"`
}

// DeleteResponse represents a response that includes errors and resource identifiers.
type DeleteResponse struct {
	Errors []Error              `json:"errors"`
	Data   []ResourceIdentifier `json:"data"`
}
