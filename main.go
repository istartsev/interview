package main

import (
	"context"
	"os"
	"syscall"
	"time"
	"os/signal"

	"github.com/sirupsen/logrus"

	"interview/websock"
)

func main() {
	ctx := context.Background()
	w := websock.NewWriter()

	go func() {
		if err := w.Run(); err != nil {
			logrus.WithError(err).Error("stopped with error")
		}
	}()

	time.Sleep(time.Second)

	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()
	h := websock.NewHandler(w, ctxCancel)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
		os.Exit(1)
	}()
	_ = h.Start("BTC/USDT")
}

