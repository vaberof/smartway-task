package pgcompany

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PgCompanyStorage struct {
	db *sqlx.DB
}

func NewPgCompanyStorage(db *sqlx.DB) *PgCompanyStorage {
	return &PgCompanyStorage{db: db}
}

func (p *PgCompanyStorage) IsExists(id int64) (bool, error) {
	query := `
			SELECT id FROM companies
			WHERE id=$1
	`
	var companyId int64
	err := p.db.QueryRow(query, id).Scan(&companyId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check whether the company exists or not: %w", err)
	}
	return true, nil
}
