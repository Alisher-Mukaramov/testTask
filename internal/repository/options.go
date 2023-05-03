package repository

import "interviewTask/pkg/database/driver/sqlx/postgres"

func WithPostgresService(p *postgres.Postgres) func(*repository) {
	return func(r *repository) {
		r.pg = p
	}
}
