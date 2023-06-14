package domain

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}
