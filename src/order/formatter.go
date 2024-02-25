package order

import (
	"reku-technical-test/src/entity"
)

type OrderFormatter struct {
	ID       int    `json:"id"`
	Pizza    string `json:"pizza"`
	Duration int    `json:"duration"`
	Status   string `json:"status"`
}

func FormatOrder(order entity.Order) OrderFormatter {
	formatOrder := OrderFormatter{
		ID:       order.ID,
		Pizza:    order.Pizza,
		Duration: order.Duration,
		Status:   order.Status,
	}

	return formatOrder
}

func FormatOrders(orders []entity.Order) []OrderFormatter {
	ordersFormatter := []OrderFormatter{}

	for _, menu := range orders {
		menuFormatter := FormatOrder(menu)
		ordersFormatter = append(ordersFormatter, menuFormatter)
	}

	return ordersFormatter
}
