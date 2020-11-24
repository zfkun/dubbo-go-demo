package game

import (
    "context"

    "github.com/apache/dubbo-go/common/logger"

    "demo/internal/common/pojo"
)

type BasketballService struct{}

func (p *BasketballService) Online(ctx context.Context, uid string) (*pojo.Result, error) {
    logger.Infof("用户上线: %s", uid)
    return &pojo.Result{Code: 0}, nil
}

func (p *BasketballService) Offline(ctx context.Context, uid string) (*pojo.Result, error) {
    logger.Infof("用户离线: %s", uid)
    return &pojo.Result{Code: 0}, nil
}

func (p *BasketballService) Message(ctx context.Context, uid, data string) (*pojo.Result, error) {
    logger.Infof("用户消息: %s, %s", uid, data)

    // 学舌回复
    _, err := gateBasketball.Send(context.TODO(), uid, data)
    if err != nil {
        logger.Errorf("推送失败: %s", err.Error())
    }

    return &pojo.Result{Code: 0}, nil
}

func (p *BasketballService) Reference() string {
    return "gameProvider.basketballService"
}


func push(uid, data string) error {
    _, err := gateBasketball.Send(context.TODO(), uid, data)
    if err != nil {
        return err
    }

    return nil
}