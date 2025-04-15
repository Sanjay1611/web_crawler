package downloader

import (
	"context"
	"testing"

	"github.com/Sanjay1611/web_crawler/pkg/result"
)

func TestURLDownloader_Download(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		d       *URLDownloader
		wantErr bool
	}{
		{
			name: "Valid URL",
			d: &URLDownloader{
				MaxConcurrent: 10,
			},
			url:     "https://www.google.com",
			wantErr: false,
		},
		{
			name: "Invalid URL",
			d: &URLDownloader{
				MaxConcurrent: 10,
			},
			url:     "google.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			in := make(chan string)
			out := make(chan result.DownloadResult)

			go tt.d.Download(ctx, in, out)
			in <- tt.url
			res := <-out

			if (res.Err != nil) != tt.wantErr {
				t.Errorf("UrlDownloader.Download() error = %v, wantErr %v", res.Err, tt.wantErr)
			}
		})
	}
}
