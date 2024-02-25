package order

type OrderRequest struct {
	PizzaID int `json:"pizza_id" binding:"required"`
}
