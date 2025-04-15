package pipelines

import (
	"context"
	"os"
	"testing"
)

func TestCsvUrlFilePipeline_Run(t *testing.T) {
	tests := []struct {
		name string
		p    Pipeline
	}{
		{
			name: "Valid Pipeline",
			p:    NewCsvUrlFilePipeline("../reader/testdata/test.csv"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Run(context.TODO())

			// check whether output directory is created
			_, err := os.Open("downloads")
			if os.IsNotExist(err) {
				t.Errorf("Output directory `downloads` is not created")
			}
		})
	}
}
