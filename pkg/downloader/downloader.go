package downloader

import (
	"context"

	"github.com/Sanjay1611/web_crawler/pkg/result"
)

type Downloader interface {
	Download(ctx context.Context, in <-chan string, out chan<- result.DownloadResult)
}
