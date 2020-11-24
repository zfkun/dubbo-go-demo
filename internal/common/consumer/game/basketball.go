package game

import (
    "context"
    "demo/internal/common/pojo"
)

type BasketballService struct {
    Online  func(ctx context.Context, uid string) (*pojo.Result, error)
    Offline func(ctx context.Context, uid string) (*pojo.Result, error)
    Message func(ctx context.Context, uid string, data string) (*pojo.Result, error)
}

func (p *BasketballService) Reference() string {
    return "gameConsumer.basketballService"
}
