package tasks

import (
	"TechSolutions/GadgetPoints/models"
	"time"
)

var OrderStatusChan = make(chan *models.Orders, 100)

func init() {
	go listenOrderStatusChan()
}

// check if order is payed or closed
func listenOrderStatusChan() {
	ticker := time.NewTicker(time.Minute)
	for {
		<-ticker.C
		for {
			order := <-OrderStatusChan
			if order == nil {
				continue
			}
			orderRecord := db.Select("state").Where("order_id=?", order.OrderID).From("order")
			if orderRecord.Status == "completed" {
				continue
			}
			if orderRecord.Status != "completed" {
				if result := db.Select("PayID").Where("order_id=?", order.OrderID).From("payment"); result == nil {
					if order.OrderDate.Sub(time.Now()) > time.Hour*24 {
						db.Update(order).Set("order_status = ?", "closed").Where("order_id = ?", order.OrderID).Exec()
						continue
					}
				} else {
					db.Update(order).Set("order_status = ?", "completed").Where("order_id = ?", order.OrderID).Exec()
					continue
				}
			}

			OrderStatusChan <- order
		}
	}
}
