package payment

import (
	"fmt"
	"sync"

	"github.com/yangpixi/GoMall/services/payment-service/internal/application/port/gateway"
	"github.com/yangpixi/GoMall/services/payment-service/internal/config"
)

type Factory func(cfg *config.Config) (gateway.PaymentGateway, error)

var (
	mu        sync.RWMutex
	factories = make(map[string]Factory)
)

type Registry struct {
	gateways map[string]gateway.PaymentGateway
}

func NewRegistry(cfg *config.Config) (*Registry, error) {
	reg := &Registry{gateways: make(map[string]gateway.PaymentGateway)}

	for channel, factory := range factories {
		g, err := factory(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to register payment gateway: %s", channel)
		}
		reg.gateways[channel] = g
	}

	return reg, nil
}

func Register(channel string, fac Factory) {
	mu.Lock()
	defer mu.Unlock()

	if _, dup := factories[channel]; dup {
		panic("duplicate registration")
	}
	factories[channel] = fac

}

func (r *Registry) Get(channel string) (gateway.PaymentGateway, error) {
	if g, ok := r.gateways[channel]; ok {
		return g, nil
	}

	return nil, fmt.Errorf("no registered gateway named: %s", channel)
}
