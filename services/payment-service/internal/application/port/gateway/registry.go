package gateway

type Resolver interface {
	Get(channel string) (PaymentGateway, error)
}
