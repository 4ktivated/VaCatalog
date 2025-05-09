package scheduler

import (
	"context"
	"some_app/pkg/parser"
	"time"

	"go.uber.org/zap"
)

type Scheduler interface {
	InitSync()
	RunSync()
	sync()
}

type ShedulerPars struct {
	logger      zap.SugaredLogger
	parseClient *parser.ParseClient
}

func NewShedilerPars(logger zap.SugaredLogger, parseClient *parser.ParseClient) *ShedulerPars {
	return &ShedulerPars{
		parseClient: parseClient,
	}
}

// TODO: подумать над планировщиком
func (s *ShedulerPars) RunSync(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(8 * time.Hour):
		}

		err := s.SyncOnce()
		if err != nil {
			s.logger.Error("can't sync stores, changes will be skipped", zap.Error(err))
		}
	}
}

func (s *ShedulerPars) SyncOnce() error {
	go func() {
		s.sync()
	}()
	return nil
}

func (s *ShedulerPars) sync() {
	parser.Processing(s.parseClient.GetPool())
}
