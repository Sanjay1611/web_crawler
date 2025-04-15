package pipelines

import (
	"context"
)

type Stage interface {
	Execute(ctx context.Context, in interface{}, out interface{})
}

// Concrete implementation with DAG

type Pipeline interface {
	Run(ctx context.Context)
}

// func (p *Pipeline) Run(ctx context.Context) {
// 	urlChan := make(chan string)
// 	downloadChan := make(chan result.DownloadResult)

// 	go p.Stages[0].Execute(ctx, nil, urlChan)
// 	go p.Stages[1].Execute(ctx, urlChan, downloadChan)
// 	p.Stages[2].Execute(ctx, downloadChan, nil)

// 	logrus.Info("Pipeline completed.")
// }
