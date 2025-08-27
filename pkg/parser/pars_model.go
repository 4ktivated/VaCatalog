package parser

import (
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

type Parser interface {
	Pars(logger *zap.SugaredLogger) *Vacs
	recvData() (*http.Response, error)
}

func Processing(parsers []Parser, logger *zap.SugaredLogger) {
	vacToRepo := make(chan Vac, 200)

	var wg sync.WaitGroup
	for _, op := range parsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vacs := op.Pars(logger)
			if len(*vacs)== 0{
				logger.Error("empty vacs")
				return
			}
			for _, vac := range *vacs {
				vacToRepo <- vac
			}

		}()
	}

	for vac := range vacToRepo {
		//TODO: put vac in db
		fmt.Printf("put vac in db: %v", vac)
	}
	wg.Wait()
}
