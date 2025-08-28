package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

//TODO: сюда вписать коннект к базе

type PgRepository struct {
	dbpool *pgxpool.Pool
	logger *zap.SugaredLogger
}

func NewRepository(dbpool *pgxpool.Pool, logger *zap.SugaredLogger) *PgRepository {
	return &PgRepository{
		dbpool: dbpool,
		logger: logger,
	}
}

func (r *PgRepository) Hello(ctx context.Context) {
	var greeting string
	err := r.dbpool.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		r.logger.Warn("QueryRow failed: %v\n", err)
	}
}
