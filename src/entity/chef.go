package entity

type Chef struct {
	ID   int    `json:"id" gorm:"primaryKey,unique"`
	Name string `json:"name"`
	Busy int    `json:"busy"`
}
