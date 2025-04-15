package downloader

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Sanjay1611/web_crawler/pkg/result"
)

type URLDownloader struct {
	MaxConcurrent int
}

func (d *URLDownloader) Download(ctx context.Context, in <-chan string, out chan<- result.DownloadResult) {
	sem := make(chan struct{}, d.MaxConcurrent)
	var wg sync.WaitGroup

	for url := range in {
		select {
		case <-ctx.Done():
			wg.Wait()
			close(out)
			return
		default:
			sem <- struct{}{}
			wg.Add(1)
			go func(url string) {
				defer func() {
					<-sem
					wg.Done()
				}()
				start := time.Now()
				resp, err := http.Get(url)
				if err != nil {
					out <- result.DownloadResult{Source: url, Err: err}
					logrus.Error(err)
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					out <- result.DownloadResult{Source: url, Err: err}
					return
				}

				logrus.Infof("Downloaded %s in %v", url, time.Since(start))
				out <- result.DownloadResult{Source: url, Content: body}
			}(url)
		}
	}
	wg.Wait()
	close(out)
}

func (s *URLDownloader) Execute(ctx context.Context, in interface{}, out interface{}) {
	urlChan, ok := in.(chan string)
	if !ok {
		logrus.Fatal("error converting interface{} to an input channel of string")
	}

	resultChan, ok := out.(chan result.DownloadResult)
	if !ok {
		logrus.Fatal("error converting interface{} to an output channel of DownloadResult")
	}

	s.Download(ctx, urlChan, resultChan)
}
