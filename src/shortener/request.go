package shortener

type ShortRequest struct {
	TargetURL string `json:"target_url" binding:"required"`
}

type FetchRequest struct {
	Shortened string `uri:"shortened" binding:"required"`
}

type GetRequest struct {
	Sort string `query:"sort"`
}
