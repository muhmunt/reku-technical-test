package chef

import (
	"reku-technical-test/src/entity"
)

type ChefFormatter struct {
	ID   int    `json:"id" gorm:"unique"`
	Name string `json:"name"`
	Busy int    `json:"busy"`
}

func FormatChef(chef entity.Chef) ChefFormatter {
	chefFormatter := ChefFormatter{
		ID:   chef.ID,
		Name: chef.Name,
		Busy: chef.Busy,
	}

	return chefFormatter
}

func FormatChefs(chefs []entity.Chef) []ChefFormatter {
	chefsFormatter := []ChefFormatter{}

	for _, chef := range chefs {
		ChefFormatter := FormatChef(chef)
		chefsFormatter = append(chefsFormatter, ChefFormatter)
	}

	return chefsFormatter
}
