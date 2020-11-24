package gate

import (
    "context"
    "demo/internal/common/pojo"
)

type GoldService struct {
    Send func(ctx context.Context, uid string, data string) (*pojo.Result, error)
}

func (p *GoldService) Reference() string {
    return "gateConsumer.goldService"
}
