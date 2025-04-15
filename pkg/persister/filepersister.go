package persister

import (
	"context"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Sanjay1611/web_crawler/pkg/result"
)

type FilePersister struct {
	OutputDir string
}

func (p *FilePersister) Persist(ctx context.Context, in <-chan result.DownloadResult) {
	_ = os.Mkdir(p.OutputDir, 0755)

	for result := range in {
		if result.Err != nil {
			logrus.WithError(result.Err).Warnf("Failed to download: %s", result.Source)
			continue
		}

		filename := filepath.Join(p.OutputDir, uuid.New().String()+".txt")
		err := os.WriteFile(filename, result.Content, 0644)
		if err != nil {
			logrus.WithError(err).Errorf("Failed to write file: %s from source %s", filename, result.Source)
			continue
		}

		logrus.Infof("Saved content from %s to %s", result.Source, filename)
	}
}

func (s *FilePersister) Execute(ctx context.Context, in interface{}, out interface{}) {
	resultChan, ok := in.(chan result.DownloadResult)
	if !ok {
		logrus.Fatal("error converting interface{} to an input channel of DownloadResult")
	}

	s.Persist(ctx, resultChan)
}
