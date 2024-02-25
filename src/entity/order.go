package entity

type Order struct {
	ID        int    `json:"id" gorm:"primaryKey,unique"`
	Pizza     string `json:"pizza"`
	Duration  int    `json:"duration"`
	Status    string    `json:"status"`
	CreatedAt int
}
