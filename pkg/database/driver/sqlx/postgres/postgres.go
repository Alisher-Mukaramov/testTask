package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const connectionPattern = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c *Config) ToString() string {
	return fmt.Sprintf(connectionPattern, c.Host, c.Port, c.User, c.Password, c.DBName)
}

type Postgres struct {
	sqlx *sqlx.DB
}

func New(cfg *Config) *Postgres {
	db, err := sqlx.Connect("postgres", cfg.ToString())
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	return &Postgres{sqlx: db}
}

func (p *Postgres) GetInstance() *sqlx.DB {
	return p.sqlx
}
