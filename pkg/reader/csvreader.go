package reader

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type CSVReader struct {
	Path string
}

func (r *CSVReader) Read(ctx context.Context, out chan<- string) error {
	file, err := os.Open(r.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil
		default:
			line := strings.TrimSpace(scanner.Text())
			lineNum++
			if lineNum == 1 {
				// expecting header "Urls"
				if strings.ToLower(line) != "urls" {
					return fmt.Errorf("invalid CSV format: expected header 'Urls' but got '%s'", line)
				}
				continue
			}

			if line != "" {
				out <- line
			}
		}
	}

	if lineNum == 0 {
		return fmt.Errorf("empty file provided as input")
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	close(out)
	return nil
}

func (s *CSVReader) Execute(ctx context.Context, in interface{}, out interface{}) {
	outChan, ok := out.(chan string)
	if !ok {
		logrus.Fatal("error converting interface{} to an output channel of string")
	}

	err := s.Read(ctx, outChan)
	if err != nil {
		logrus.WithError(err).Fatal("Reader error")
	}
}
