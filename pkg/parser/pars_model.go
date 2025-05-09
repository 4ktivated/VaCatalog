package parser

import "context"

type Parser[P any] interface {
	Pars()
}

func Processing[S any](ctx context.Context, ops ...Executable[S]) error {
	for _, op := range ops {
		// do parse stuff
	}

	return nil
}
