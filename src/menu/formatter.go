package menu

import (
	"reku-technical-test/src/entity"
)

type MenuFormatter struct {
	ID       int    `json:"id" gorm:"unique"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

func FormatMenu(menu entity.Menu) MenuFormatter {
	formatMenus := MenuFormatter{
		ID:       menu.ID,
		Name:     menu.Name,
		Duration: menu.Duration,
	}

	return formatMenus
}

func FormatMenus(menus []entity.Menu) []MenuFormatter {
	menusFormatter := []MenuFormatter{}

	for _, menu := range menus {
		menuFormatter := FormatMenu(menu)
		menusFormatter = append(menusFormatter, menuFormatter)
	}

	return menusFormatter
}
