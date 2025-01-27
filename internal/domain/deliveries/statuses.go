package deliveries

type DeliveryStatus string

const (
	DeliveryStatusScheduled DeliveryStatus = "запланирована"
	DeliveryStatusOnTheWay  DeliveryStatus = "в пути"
	DeliveryStatusCanceled  DeliveryStatus = "отменена"
	DeliveryStatusCompleted DeliveryStatus = "завершена"
)
