package game

import (
    "demo/internal/common/pojo"
    hessian "github.com/apache/dubbo-go-hessian2"
    "github.com/apache/dubbo-go/config"

    _ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
    _ "github.com/apache/dubbo-go/protocol/dubbo"
    _ "github.com/apache/dubbo-go/registry/protocol"

    _ "github.com/apache/dubbo-go/filter/filter_impl"
    _ "github.com/apache/dubbo-go/filter/filter_impl/auth"

    _ "github.com/apache/dubbo-go/cluster/cluster_impl"
    _ "github.com/apache/dubbo-go/cluster/loadbalance"
    // _ "github.com/apache/dubbo-go/metadata/report/zookeeper"
    // _ "github.com/apache/dubbo-go/metadata/service/remote"
    // _ "github.com/apache/dubbo-go/metadata/service/inmemory"
    _ "github.com/apache/dubbo-go/registry/zookeeper"
)

func Start() {
    config.SetProviderService(new(BasketballService))
    config.SetConsumerService(gateBasketball)

    hessian.RegisterPOJO(&pojo.Result{})

    config.Load()
}
