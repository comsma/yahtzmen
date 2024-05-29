package ports

type PaymentPort interface {
	CreatePaymentIntent(amount float64, currency string) (*string, error)
}
