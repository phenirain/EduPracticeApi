package payments

type PaymentStatus string

const (
	PaymentStatusWaiting  PaymentStatus = "ожидает обработки"
	PaymentStatusCanceled PaymentStatus = "отменен"
	PaymentStatusPayed    PaymentStatus = "оплачен"
)
