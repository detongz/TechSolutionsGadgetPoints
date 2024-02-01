package api

import "TechSolutions/GadgetPoints/models"

// 用户下单
// GetUserStatus retrieves the subscription status, active discounts, and current points balance for a user
func GetUserStatus(userID int) (isSubscriber bool, discounts []Discount, points int, err error) {
	return
}

// CreateOrder creates a new order for the given user at the specified outlet, with the provided products and accessories
func CreateOrder(userID int, outletID int, products []int, accessories []int) (Order, error) {
	discounts := models.QueryDiscounts(outletID)

	subscribe := models.QueryUserSubscribe(userID)

	cost := CalculateCost(discounts, subscribe)

	models.AddOrder(models.Orders{
		Products:    products,
		Accessories: accessories,
		OutletID:    outletID,
		UserID:      userID,
		Discounts:   discounts,
		TotalCost:   cost,
	})

	return
}

// CancelOrder cancels an existing order and credits the points back to the user
func CancelOrder(orderID int) error {
	// Logic to cancel an order and credit points back to the user's account
	return nil
}

// CompleteOrder marks an order as completed and adds the earned points to the user's account
func CompleteOrder(orderID int, totalCost float64) error {
	// 查询支付情况
	// 如果查询支付情况属实，那么给用户计算积分
	return nil
}
