package gate

import (
    "context"
    "demo/internal/common/pojo"
)

type BasketballService struct {
    Send func(ctx context.Context, uid string, data string) (*pojo.Result, error)
}

func (p *BasketballService) Reference() string {
    return "gateConsumer.basketballService"
}
