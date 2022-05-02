package main

import (
	"os"
	"os/signal"
	"syscall"
	_ "vsign/api"
	"vsign/cleanup"
	"vsign/logger"
)

func main() {
	logger.Info("checking env......")

	// exit 监听
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case s := <-c:
		cleanup.ExitService(s)
	}
}
