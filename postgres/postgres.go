package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

type DBLogger struct {}

func (p DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (p DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) (error) {
	fmt.Println(q.FormattedQuery())

	return nil
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}