package pgemployee

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vaberof/smartway-task/internal/domain/employee"
	"github.com/vaberof/smartway-task/internal/infra/storage"
)

const (
	errorEmployeeId   = 0
	errorDepartmentId = 0
)

type PgEmployeeStorage struct {
	db *sqlx.DB
}

func NewPgEmployeeStorage(db *sqlx.DB) *PgEmployeeStorage {
	return &PgEmployeeStorage{db: db}
}

func (p *PgEmployeeStorage) Create(name, surname, phone string, companyId int64, passportType, passportNumber string, departmentName, departmentPhone string) (int64, error) {
	query := `
			INSERT INTO employees(
			                      name, 
			                      surname,
			                      phone,
			                      company_id,
			                      passport,
			                      department
			) VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id
	`

	pgPassport := &Passport{
		Type:   passportType,
		Number: passportNumber,
	}

	pgPassportBytes, _ := pgPassport.Value()

	pgDepartment := &Department{
		Name:  departmentName,
		Phone: departmentPhone,
	}

	pgDepartmentBytes, _ := pgDepartment.Value()

	row := p.db.QueryRow(query, name, surname, phone, companyId, pgPassportBytes, pgDepartmentBytes)

	var employeeId int64

	err := row.Scan(&employeeId)
	if err != nil {
		return errorEmployeeId, fmt.Errorf("failed to create an employee: %w", err)
	}

	return employeeId, nil
}

func (p *PgEmployeeStorage) Update(id int64, name, surname, phone *string, companyId *int64, passportType, passportNumber, departmentName, departmentPhone *string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction while updating an employee: %w", err)
	}
	defer tx.Rollback()

	query := `
			UPDATE employees
						SET 
							name = COALESCE($1, name),
							surname = COALESCE($2, surname),
							phone = COALESCE($3, phone),
							company_id = COALESCE($4, company_id),
							passport = COALESCE($5, passport),
							department = COALESCE($6, department)
						WHERE id=$7
	`

	var pgPassport *Passport

	if passportType != nil || passportNumber != nil {
		pgPassport, err = p.getUpdatedPassport(tx, id, passportType, passportNumber)
		if err != nil {
			return fmt.Errorf("failed to update an employee: %w", err)
		}
	}

	var pgDepartment *Department

	if departmentName != nil || departmentPhone != nil {
		pgDepartment, err = p.getUpdatedDepartment(tx, id, departmentName, departmentPhone)
		if err != nil {
			return fmt.Errorf("failed to update an employee: %w", err)
		}
	}

	pgPassportBytes, _ := pgPassport.Value()
	pgDepartmentBytes, _ := pgDepartment.Value()

	result, err := tx.Exec(query, name, surname, phone, companyId, pgPassportBytes, pgDepartmentBytes, id)
	if err != nil {
		return fmt.Errorf("failed to update an employee: %w", err)
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("failed to update an employee: %w", storage.ErrEmployeeNotFound)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction while updating an employee: %w", err)
	}

	return nil
}

func (p *PgEmployeeStorage) getUpdatedPassport(tx *sql.Tx, employeeId int64, passportType, passportNumber *string) (*Passport, error) {
	queryGetPassport := `
					SELECT passport FROM employees WHERE id=$1
`
	var pgPassport *Passport
	var passportBytes []byte

	row := tx.QueryRow(queryGetPassport, employeeId)
	if err := row.Scan(&passportBytes); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEmployeeNotFound
		}

		return nil, err
	}

	if len(passportBytes) > 0 {
		if err := json.Unmarshal(passportBytes, &pgPassport); err != nil {
			return nil, err
		}
	}

	if passportType != nil {
		pgPassport.Type = *passportType
	}
	if passportNumber != nil {
		pgPassport.Number = *passportNumber
	}

	return pgPassport, nil
}

func (p *PgEmployeeStorage) getUpdatedDepartment(tx *sql.Tx, employeeId int64, departmentName, departmentPhone *string) (*Department, error) {
	queryGetDepartment := `
					SELECT department FROM employees WHERE id=$1
`
	var pgDepartment *Department
	var departmentBytes []byte

	row := tx.QueryRow(queryGetDepartment, employeeId)
	if err := row.Scan(&departmentBytes); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEmployeeNotFound
		}

		return nil, err
	}

	if len(departmentBytes) > 0 {
		if err := json.Unmarshal(departmentBytes, &pgDepartment); err != nil {
			return nil, err
		}
	}

	if departmentName != nil {
		pgDepartment.Name = *departmentName
	}
	if departmentPhone != nil {
		pgDepartment.Phone = *departmentPhone
	}

	return pgDepartment, nil
}

func (p *PgEmployeeStorage) Delete(id int64) error {
	query := `DELETE FROM employees WHERE id=$1`
	result, err := p.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete an employee: %w", err)
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("failed to delete an employee: %w", storage.ErrEmployeeNotFound)
	}
	return nil
}

func (p *PgEmployeeStorage) ListByCompanyId(companyId int64, limit, offset int) ([]*employee.Employee, error) {
	limitOffsetParams := fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)

	query := `
			SELECT * FROM employees WHERE company_id=$1 
`
	query += limitOffsetParams

	rows, err := p.db.Query(query, companyId)
	if err != nil {
		return nil, fmt.Errorf("failed to list employees by company id: %w", err)
	}
	defer rows.Close()

	var pgEmployees []*PgEmployee

	for rows.Next() {
		var pgEmployee PgEmployee

		err = rows.Scan(
			&pgEmployee.Id,
			&pgEmployee.Name,
			&pgEmployee.Surname,
			&pgEmployee.Phone,
			&pgEmployee.CompanyId,
			&pgEmployee.Passport,
			&pgEmployee.Department,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to list employees by company id: %w", err)
		}

		pgEmployees = append(pgEmployees, &pgEmployee)
	}

	return p.buildDomainEmployees(pgEmployees), nil
}

func (p *PgEmployeeStorage) ListByDepartmentName(companyId int64, departmentName string, limit, offset int) ([]*employee.Employee, error) {
	limitOffsetParams := fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)

	query := `
			SELECT * FROM employees WHERE company_id=$1 AND department->>'name'=$2
`
	query += limitOffsetParams

	rows, err := p.db.Query(query, companyId, departmentName)
	if err != nil {
		return nil, fmt.Errorf("failed to list employees by department name: %w", err)
	}
	defer rows.Close()

	var pgEmployees []*PgEmployee

	for rows.Next() {
		var pgEmployee PgEmployee

		err = rows.Scan(
			&pgEmployee.Id,
			&pgEmployee.Name,
			&pgEmployee.Surname,
			&pgEmployee.Phone,
			&pgEmployee.CompanyId,
			&pgEmployee.Passport,
			&pgEmployee.Department,
		)
		if err != nil {

			return nil, fmt.Errorf("failed to list employees by department name: %w", err)
		}

		pgEmployees = append(pgEmployees, &pgEmployee)
	}

	return p.buildDomainEmployees(pgEmployees), nil
}

func (p *PgEmployeeStorage) buildDomainEmployee(pgEmployee *PgEmployee) *employee.Employee {
	return &employee.Employee{
		Id:        pgEmployee.Id,
		Name:      pgEmployee.Name,
		Surname:   pgEmployee.Surname,
		Phone:     pgEmployee.Phone,
		CompanyId: pgEmployee.CompanyId,
		Passport: employee.Passport{
			Type:   pgEmployee.Passport.Type,
			Number: pgEmployee.Passport.Number,
		},
		Department: employee.Department{
			Name:  pgEmployee.Department.Name,
			Phone: pgEmployee.Department.Phone,
		},
	}
}

func (p *PgEmployeeStorage) buildDomainEmployees(pgEmployees []*PgEmployee) []*employee.Employee {
	domainEmployees := make([]*employee.Employee, len(pgEmployees))
	for i := range pgEmployees {
		domainEmployees[i] = p.buildDomainEmployee(pgEmployees[i])
	}
	return domainEmployees
}
