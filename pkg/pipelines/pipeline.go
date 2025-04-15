package pipelines

import (
	"context"
	"sync"
)

type Stage interface {
	Execute(ctx context.Context, in interface{}, out interface{})
}

type Pipeline interface {
	Run(ctx context.Context)
}

type pipeline struct {
	stages []Stage
	wg     *sync.WaitGroup
}

// func (p *Pipeline) Run(ctx context.Context) {
// 	urlChan := make(chan string)
// 	downloadChan := make(chan result.DownloadResult)

// 	go p.Stages[0].Execute(ctx, nil, urlChan)
// 	go p.Stages[1].Execute(ctx, urlChan, downloadChan)
// 	p.Stages[2].Execute(ctx, downloadChan, nil)

// 	logrus.Info("Pipeline completed.")
// }
