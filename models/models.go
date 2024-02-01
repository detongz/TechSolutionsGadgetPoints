package models

import (
	"time"
)

// Users struct represents the Users table
type Users struct {
	UserID   int       // Unique identifier for the user
	Username string    // User's chosen username
	Password string    // User's hashed password
	Email    string    // User's email address
	JoinDate time.Time // Timestamp of when the user joined
}

// Products struct represents the Products table
type Products struct {
	OutletID      int
	ProductID     int     // Unique identifier for the product
	ProductName   string  // Name of the product
	Description   string  // Description of the product
	Price         float64 // Price of the product
	StockQuantity int     // Quantity of the product in stock
}

// Accessories struct represents the Accessories table
type Accessories struct {
	OutletID      int
	ProductID     int
	AccessoryID   int     // Unique identifier for the accessory
	AccessoryName string  // Name of the accessory
	Description   string  // Description of the accessory
	Price         float64 // Price of the accessory
	StockQuantity int     // Quantity of the accessory in stock
}

// Outlets struct represents the Outlets table
type Outlets struct {
	OutletID     int    // Unique identifier for the outlet
	OutletName   string // Name of the outlet
	Address      string // Address of the outlet
	Phone        string // Phone number of the outlet
	OpeningHours string // Opening hours of the outlet
}

// Orders struct represents the Orders table
type Orders struct {
	OrderID     int // Unique identifier for the order
	UserID      int // Foreign key referencing Users.UserID
	OutletID    int // Foreign key referencing Outlets.OutletID
	Products    []int
	Accessories []int
	TotalCost   float64    // Total cost of the order
	Discounts   []Discount // List of discounts applied to the order
	OrderDate   time.Time  // Timestamp of when the order was placed
	Status      string     // Status of the order (e.g., Pending, Completed, Cancelled)
}

// 直接减价格、折扣、送积分等等，Amount是对应的数量
type Discount struct {
	DiscountID int     // Unique identifier for the discount
	Type       string  // Type of discount (e.g., loyalty, subscription)
	Range      string  // national or Outlet
	Amount     float64 // Amount of the discount
}

type Payment struct {
	PayID   int
	OrderID int // Foreign key referencing Orders.OrderID
	Amount  float64
	Date    time.Time
}

// LoyaltyRedemptions struct represents the LoyaltyRedemptions table
type LoyaltyRedemptions struct {
	RedemptionID   int       // Unique identifier for the redemption
	UserID         int       // Foreign key referencing Users.UserID
	AccessoryID    int       // Foreign key referencing Accessories.AccessoryID
	RedemptionDate time.Time // Timestamp of when the redemption was made
}

// Subscriptions struct represents the Subscriptions table
type Subscriptions struct {
	SubscriptionID   int       // Unique identifier for the subscription
	UserID           int       // Foreign key referencing Users.UserID
	SubscriptionType string    // Type of the subscription
	StartDate        time.Time // Start date of the subscription
	EndDate          time.Time // End date of the subscription
}

// ServicingRecords struct represents the ServicingRecords table
type ServicingRecords struct {
	ServiceRecordID    int       // Unique identifier for the service record
	SubscriptionID     int       // Foreign key referencing Subscriptions.SubscriptionID
	DeviceID           int       // Unique identifier for the device
	ServiceDate        time.Time // Timestamp of when the service was performed
	ServiceDescription string    // Description of the service performed
}
