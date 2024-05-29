package payment

import (
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/client"
	"net/http"
)

type Adapter struct {
	stripeClient *client.API
}

func NewAdapter() *Adapter {
	return &Adapter{
		stripeClient: client.New("", stripe.NewBackends(http.DefaultClient)),
	}
}

func (a *Adapter) CreatePaymentIntent(amount float64, currency string) (*string, error) {

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount * 100)),
		Currency: stripe.String(currency),
	}
	pi, err := a.stripeClient.PaymentIntents.New(params)
	if err != nil {
		return nil, err
	}
	return &pi.ClientSecret, nil
}
