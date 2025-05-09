package parser

import (
	"fmt"
	"net/http"
	"sync"
)

type Parser interface {
	Pars() (*Vacs, error)
	recvData() (*http.Response, error)
}

// type Service[A any, S any] interface {
// 	Run(ctx context.Context) (*Vacs, error)
// }
//
// type Operation[A any, S any] struct {
// 	Repository repository.Data
// 	Service    Service[A, S]
// }

// func (o *Operation[A, S]) Pars(ctx context.Context) (*Vacs, error) {
// 	vacs, err := o.Service.Run(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("cant run process parser: %w", err)
// 	}
// 	return vacs, nil
// }

func Processing(parsers []Parser) {
	vacToRepo := make(chan Vac, 200)

	var wg sync.WaitGroup
	for _, op := range parsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vacs, err := op.Pars()
			if err != nil {
				return
			}
			for _, vac := range *vacs {
				vacToRepo <- vac
			}

		}()
	}

	for vac := range vacToRepo {
		fmt.Printf("put vac in db: %v", vac)
	}
	wg.Wait()
}
