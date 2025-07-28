package store

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PG struct {
	db *sql.DB
}

func MustPostgres(dsn string) *PG {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	return &PG{db: db}
}

// Example: сохранить снапшот -------------------------------------------------

func (p *PG) SaveSnapshot(ctx context.Context, epoch uint64, blob []byte) error {
	_, err := p.db.ExecContext(
		ctx,
		`INSERT INTO snapshots(epoch, data) VALUES ($1, $2)
         ON CONFLICT (epoch) DO NOTHING`,
		epoch, blob,
	)
	return err
}
