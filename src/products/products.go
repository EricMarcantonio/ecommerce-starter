package products

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Desc    string  `json:"desc,omitempty"`
	Price   float64 `json:"price"`
	TakenBy string  `json:"taken_by"`
}
