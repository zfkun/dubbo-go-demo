package gate

import (
    "context"
    "demo/internal/common/pojo"
    "net/http"

    hessian "github.com/apache/dubbo-go-hessian2"
    "github.com/apache/dubbo-go/config"
    _ "github.com/apache/dubbo-go/protocol/dubbo"
    _ "github.com/apache/dubbo-go/registry/protocol"

    _ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
    _ "github.com/apache/dubbo-go/filter/filter_impl"

    _ "github.com/apache/dubbo-go/cluster/cluster_impl"
    _ "github.com/apache/dubbo-go/cluster/loadbalance"
    // _ "github.com/apache/dubbo-go/metadata/report/zookeeper"
    // _ "github.com/apache/dubbo-go/metadata/service/remote"
    _ "github.com/apache/dubbo-go/metadata/service/inmemory"
    _ "github.com/apache/dubbo-go/registry/zookeeper"
)

func Start() {
    config.SetProviderService(new(BasketballService)) // 激活
    config.SetProviderService(new(JumpService))       // 激活
    config.SetProviderService(new(GoldService))       // 不激活

    config.SetConsumerService(gameBasketball)

    hessian.RegisterPOJO(&pojo.Result{})

    config.Load()

    go startHttp()
}

func startHttp() {
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        gameBasketball.Message(context.TODO(), "abc", "测试消息")
        w.Write([]byte("ok"))
    })
    http.ListenAndServe(":8000", nil)
}
