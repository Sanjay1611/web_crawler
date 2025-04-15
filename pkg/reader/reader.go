package reader

import "context"

type Reader interface {
	Read(ctx context.Context, out chan<- string) error
}
