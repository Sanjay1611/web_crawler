package pipelines

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Sanjay1611/web_crawler/pkg/downloader"
	"github.com/Sanjay1611/web_crawler/pkg/persister"
	"github.com/Sanjay1611/web_crawler/pkg/reader"
	"github.com/Sanjay1611/web_crawler/pkg/result"
)

type CsvUrlFilePipeline struct {
	Stages []Stage
}

func NewCsvUrlFilePipeline(csvPath string) Pipeline {
	reader := &reader.CSVReader{Path: csvPath}
	downloader := &downloader.URLDownloader{MaxConcurrent: 50}
	persister := &persister.FilePersister{OutputDir: "downloads"}

	pipeline := &CsvUrlFilePipeline{
		Stages: []Stage{reader, downloader, persister},
	}

	return pipeline
}

func (p *CsvUrlFilePipeline) Run(ctx context.Context) {
	urlChan := make(chan string)
	downloadChan := make(chan result.DownloadResult)

	go p.Stages[0].Execute(ctx, nil, urlChan)
	go p.Stages[1].Execute(ctx, urlChan, downloadChan)
	p.Stages[2].Execute(ctx, downloadChan, nil)

	logrus.Info("Pipeline completed.")
}
