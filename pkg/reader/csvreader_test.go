package reader

import (
	"context"
	"testing"
)

func TestCSVReader_Read(t *testing.T) {
	tests := []struct {
		name    string
		r       *CSVReader
		wantErr bool
	}{
		{
			name: "empty file",
			r: &CSVReader{
				Path: "testdata/empty.csv",
			},
			wantErr: true,
		},
		{
			name: "file without header",
			r: &CSVReader{
				Path: "testdata/withoutHeader.csv",
			},
			wantErr: true,
		},
		{
			name: "test file",
			r: &CSVReader{
				Path: "testdata/test.csv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			out := make(chan string)

			// read from the channel so that it does not get blocked
			go func() {
				for {
					if _, ok := <-out; !ok {
						return
					}
				}
			}()
			if err := tt.r.Read(ctx, out); (err != nil) != tt.wantErr {
				t.Errorf("CSVReader.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
