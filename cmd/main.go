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
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	if len(os.Args) < 2 {
		logrus.Fatal("Usage: go run cmd/main.go <path_to_csv>")
	}

	ctx, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(stop, cancel)

	// assuming os.Args[1] to be csv file path to be read
	pipes := []pipelines.Pipeline{pipelines.NewCsvUrlFilePipeline(os.Args[1])}
	pipelines.NewPipelineRunner(pipes).Run(ctx)
}

func gracefulShutdown(stop chan os.Signal, cancel context.CancelFunc) {
	<-stop
	logrus.Info("Interrupt received, initiating shutdown...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	cancel()

	<-shutdownCtx.Done()
	logrus.Info("Shutdown time reached.")
}
