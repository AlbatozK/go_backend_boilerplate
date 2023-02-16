package repository

import (
	"errors"

	"github.com/AlbatozK/go_backend_boilerplate/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type GenericRepository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewGenericRepository() *GenericRepository {
	return &GenericRepository{}
}

func (gr *GenericRepository) BeginTx() error {
	if gr.db == nil {
		gr.db = db.GetPool()
	}
	tx, err := gr.db.Beginx()
	if err != nil {
		return err
	}
	gr.tx = tx
	return nil
}

func (gr *GenericRepository) Rollback() error {
	if gr.tx == nil {
		return errors.New("transaction not started")
	}
	err := gr.tx.Rollback()
	gr.tx = nil
	return err
}

func (gr *GenericRepository) Commit() error {
	if gr.tx == nil {
		return errors.New("transaction not started")
	}
	err := gr.tx.Commit()
	gr.tx = nil
	return err
}

func (gr *GenericRepository) NewRepository() *GenericRepository {
	return &GenericRepository{tx: gr.tx}
}
