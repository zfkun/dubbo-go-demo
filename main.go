package main

import (
    "demo/internal/game"
    "demo/internal/gate"
    "flag"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/apache/dubbo-go/common/logger"
)

var serverType string

func main() {
    flag.StringVar(&serverType, "t", "gate", "start gate server")
    flag.Parse()

    if serverType == "gate" {
        logger.Infof("start gate server")
        gate.Start()
    } else if serverType == "game" {
        logger.Infof("start game server")
        game.Start()
    } else {
        panic("server type invalid")
    }

    initSignal()
}

func initSignal() {
    signals := make(chan os.Signal, 1)

    signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
    for {
        sig := <-signals
        logger.Infof("get signal %s", sig.String())
        switch sig {
        case syscall.SIGHUP:
            logger.Infof("app need reload")
        default:
            time.AfterFunc(time.Duration(time.Second*5), func() {
                logger.Warnf("app exit now by force...")
                os.Exit(1)
            })

            // The program exits normally or timeout forcibly exits.
            logger.Warnf("app exit now...")
            return
        }
    }
}
