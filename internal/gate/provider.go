package gate

import (
    "context"
    "github.com/apache/dubbo-go/common/logger"

    "demo/internal/common/pojo"
)

type BasketballService struct{}

func (p *BasketballService) Send(ctx context.Context, uid, data string) (*pojo.Result, error) {
    logger.Infof("basketball消息: %s", uid, data)
    return &pojo.Result{Code: 0}, nil
}

func (p *BasketballService) Reference() string {
    return "gateProvider.basketballService"
}

type JumpService struct{}

func (p *JumpService) Send(ctx context.Context, uid, data string) (*pojo.Result, error) {
    logger.Infof("jump消息: %s, %s", uid, data)
    return &pojo.Result{Code: 0}, nil
}

func (p *JumpService) Reference() string {
    return "gateProvider.jumpService"
}

type GoldService struct{}

func (p *GoldService) Send(ctx context.Context, uid, data string) (*pojo.Result, error) {
    logger.Infof("gold消息: %s, %s", uid, data)
    return &pojo.Result{Code: 0}, nil
}

func (p *GoldService) Reference() string {
    return "gateProvider.goldService"
}
