package persister

import (
	"context"

	"github.com/Sanjay1611/web_crawler/pkg/result"
)

type Persister interface {
	Persist(ctx context.Context, in <-chan result.DownloadResult)
}
