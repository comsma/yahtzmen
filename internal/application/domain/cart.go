package domain

type Cart struct {
	Products []*CartItem
}

type CartItem struct {
	ProductId string
	Quantity  int
}
