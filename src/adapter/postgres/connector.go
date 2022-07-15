package postgres

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx" //driver pgx used to run migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"log"
)

// PoolInterface is an wrapping to PgxPool to creat test mocks
type PoolInterface interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(
		ctx context.Context,
		sql string,
		args []interface{},
		scans []interface{},
		f func(pgx.QueryFuncRow) error,
	) (pgconn.CommandTag, error)
	SendBatch(tx context.Context, b *pgx.Batch) pgx.BatchResults
	//Begin(ctx context.Context) (*pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(tx pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(tx pgx.Tx) error) error
}

// GetConnection return an connection pool from postgres drive PGX
func GetConnection(context context.Context) *pgxpool.Pool {
	databaseURL := viper.GetString("database.url")

	conn, err := pgxpool.Connect(context, "postgres"+databaseURL)
	if err != nil {
		log.Panicf("Unable to connect to database: %v", err)
	}

	return conn
}

// RunMigrations run scripts on path src/database/migrations
func RunMigrations() {
	databaseUrl := viper.GetString("database.url")
	m, err := migrate.New("file://src/database/migrations", "pgx"+databaseUrl)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
