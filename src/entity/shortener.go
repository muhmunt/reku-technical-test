package entity

type Shortener struct {
	ID         string `json:"id" gorm:"primaryKey,unique"`
	TargetURL  string `json:"target_url"`
	ExpiryDate int    `json:"expiry_date"`
	Clicks     int    `json:"clicks"`
	CreatedAt  int
	UpdatedAt  int
}
