package dto

type AdminDashboardResponse struct {
	TotalPaintings     int64   `json:"total_paintings"`
	AvailablePaintings int64   `json:"available_paintings"`
	SoldPaintings      int64   `json:"sold_paintings"`
	FeaturedPaintings  int64   `json:"featured_paintings"`

	TotalOrders     int64 `json:"total_orders"`
	PendingOrders   int64 `json:"pending_orders"`
	ConfirmedOrders int64 `json:"confirmed_orders"`
	ShippedOrders   int64 `json:"shipped_orders"`
	DeliveredOrders int64 `json:"delivered_orders"`
	CancelledOrders int64 `json:"cancelled_orders"`

	TotalCustomers int64 `json:"total_customers"`

	TotalRevenue float64 `json:"total_revenue"`
}