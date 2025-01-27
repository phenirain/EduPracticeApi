package deliveries

type DeliveryStatus string

const (
	DeliveryStatusScheduled DeliveryStatus = "Scheduled"
	DeliveryStatusOnTheWay  DeliveryStatus = "OnTheWay"
	DeliveryStatusCanceled  DeliveryStatus = "Canceled"
	DeliveryStatusCompleted DeliveryStatus = "Completed"
)
