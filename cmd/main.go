package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sanjay1611/web_crawler/pkg/pipelines"
	"github.com/sirupsen/logrus"
)

// Main Execution

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// Graceful shutdown on interrupt
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(sigs, cancel)

	if len(os.Args) < 2 {
		logrus.Fatal("Usage: go run main.go <path_to_csv>")
	}

	// assuming os.Args[1] to be csv file path to be read
	pipelines.NewCsvUrlFilePipeline(os.Args[1]).Run(ctx)
}

func gracefulShutdown(sigs chan os.Signal, cancel context.CancelFunc) {
	<-sigs
	logrus.Info("Interrupt received, initiating shutdown...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	cancel()
	<-shutdownCtx.Done()
}
