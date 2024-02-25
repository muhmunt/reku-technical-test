package chef

type CheftRequest struct {
	Name string `json:"name" binding:"required"`
}
