package entity

type Menu struct {
	ID       int    `json:"id" gorm:"primaryKey,unique"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}
