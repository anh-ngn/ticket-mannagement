package main

import (
	"go.tekoapis.com/tekone/app/warehouse/iam_service/config"
	"go.tekoapis.com/tekone/app/warehouse/iam_service/internal/service"
	"go.tekoapis.com/tekone/app/warehouse/iam_service/internal/store"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func newDB(dsn string) (*sqlx.DB, error) {
	/**
	|-------------------------------------------------------------------------
	| Code generated by tekit-gen-project with database option.
	| @notice: exactly the same as the built-in
	|-----------------------------------------------------------------------*/
	db, err := sqlx.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}
	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newService(cfg *config.Config) (*service.Service, error) {
	// Code generated by tekit-gen-project with database option.
	// log cfg.PostgreSQL.DSN()
	db, err := newDB(cfg.PostgreSQL.String())

	if err != nil {
		logger.Error(err, "Error connect database")
		return nil, err
	}

	store := store.NewStore(db)

	// more adapter here

	return service.NewService(logger, store), nil
}
