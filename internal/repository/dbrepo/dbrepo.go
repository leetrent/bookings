package dbrepo

import (
	"database/sql"

	"github.com/leetrent/bookings/internal/config"
	"github.com/leetrent/bookings/internal/repository"
)

type postgresDBRepo struct {
	AppConfig *config.AppConfig
	DB        *sql.DB
}

func NewPostgresRep(conn *sql.DB, ac *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		AppConfig: ac,
		DB:        conn,
	}
}
