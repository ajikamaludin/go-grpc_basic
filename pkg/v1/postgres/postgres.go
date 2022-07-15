package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/config"
	_ "github.com/lib/pq"
)

const (
	POSTGRES string = "postgres"
)

type Conn struct {
	Conn string
}

func New(config *config.Config) (*Conn, error) {
	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Username,
		config.Postgres.Password,
		config.Postgres.Dbname)
	db, err := sql.Open(POSTGRES, conn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	return &Conn{
		Conn: conn,
	}, nil
}
