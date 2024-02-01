package api

import "TechSolutions/GadgetPoints/models"

func CalculateCost(discounts []models.Discount, subscribe models.Subscriptions, products []models.Products) (cost float64) {
	for p := range products {
		for d := range discounts {
			cost += discounts(d, p.Price)
		}
	}

	return
}

func discount(discount models.Discount, price float64) (cost float64) {
	switch discount.Type {
	case "percentage":
		//...
	case "points":
	// ...
	case "fixed":
		//...
	}
	return
}
