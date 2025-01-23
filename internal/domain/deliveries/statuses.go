package deliveries

type DeliveryStatus string

const (
	DeliveryStatusSchedule  DeliveryStatus = "запланирован"
	DeliveryStatusActive    DeliveryStatus = "активен"
	DeliveryStatusRoute     DeliveryStatus = "в пути"
	DeliveryStatusCanceled  DeliveryStatus = "отменен"
	DeliveryStatusCompleted DeliveryStatus = "завершен"
)
