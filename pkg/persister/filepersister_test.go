package persister

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Sanjay1611/web_crawler/pkg/result"
)

func TestFilePersister_Persist(t *testing.T) {
	tests := []struct {
		name string
		p    *FilePersister
		d    *result.DownloadResult
	}{
		{
			name: "Valid result",
			p: &FilePersister{
				OutputDir: "tmp",
			},
			d: &result.DownloadResult{
				Source:  "validUrl",
				Content: []byte("some content"),
				Err:     nil,
			},
		},
		{
			name: "Error result",
			p: &FilePersister{
				OutputDir: "tmp",
			},
			d: &result.DownloadResult{
				Source:  "invalidUrl",
				Content: []byte("some content"),
				Err:     fmt.Errorf("some error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			in := make(chan result.DownloadResult)

			go tt.p.Persist(ctx, in)
			in <- *tt.d

			// check whether output directory is created
			dir, err := os.Open(tt.p.OutputDir)
			if os.IsNotExist(err) {
				t.Errorf("Output directory %v is not created", tt.p.OutputDir)
			}

			// check whether output directory has file created for valid download result
			_, err = dir.Readdir(1)
			if tt.d.Err == nil && err == io.EOF {
				t.Errorf("No output file created for source %s", tt.d.Source)
			} else if tt.d.Err != nil && err != io.EOF {
				t.Errorf("Output file should not be created for source %s", tt.d.Source)
			}

			// remove output dir
			dir.Close()
			err = os.RemoveAll(tt.p.OutputDir)
			if err != nil {
				t.Errorf("why this err %v", err)
			}
		})
	}
}
