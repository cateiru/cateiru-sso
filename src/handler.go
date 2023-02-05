package src

import "database/sql"

type Handler struct {
	DB *sql.DB
	C  *Config
}

func NewHandler(db *sql.DB, config *Config) *Handler {
	return &Handler{
		DB: db,
		C:  config,
	}
}
