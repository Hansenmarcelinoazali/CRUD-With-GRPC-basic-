package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"latihan/pkg/v1/config"
)

const (
	POSTGRES string = "postgres"
)

type Conn struct {
	Conn string
}

// New initializes the postgres writers
func New(config *config.Config) (*Conn, error) {
	// Set postgres och config
	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Pass,
		config.Postgres.Dbname)

	_, err := sql.Open(POSTGRES, conn)
	if err != nil {
		return nil, err

	}
	return &Conn{
		Conn: conn,
	}, nil

}
