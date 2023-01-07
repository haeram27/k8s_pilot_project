package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"

	"pilot_agent/apps"
	"pilot_agent/routers"
)

func runEchoServer(ctx context.Context) {
	done := make(chan bool)

	echo := apps.GetEcho()
	echo.Logger.SetLevel(log.DEBUG)

	err := apps.ReadConfig()
	if err != nil {
		apps.Logs.Fatal(err.Error())
	}

	routers.Route(echo)

	go func() {
		err := echo.Start(":" + apps.Conf.Agent.Port)
		if err != nil && err != http.ErrServerClosed {
			apps.Logs.Fatal(err.Error())
		}

		done <- true
	}()

	select {
	case <-ctx.Done():
		apps.Logs.Info("cancel")
	case <-done:
		apps.Logs.Info("done")
	}
}

func main() {
	event := make(chan os.Signal)
	signal.Notify(event, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		apps.Logs.Info("signal:", <-event)
		cancel()
	}()

	runEchoServer(ctx)
}
