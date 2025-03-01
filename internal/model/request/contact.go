package request

type CreateContact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Notes string `json:"notes"`
}

type UpdateContact struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Notes *string `json:"notes,omitempty"`
}
