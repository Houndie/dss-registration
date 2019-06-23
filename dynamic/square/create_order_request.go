package square

type CreateOrderRequest struct {
	Order          *Order `json:"order"`
	IdempotencyKey string `json:"idempotency_key"`
}
