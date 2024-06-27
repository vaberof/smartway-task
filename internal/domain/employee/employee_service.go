package employee

import (
	"errors"
	"fmt"
	"github.com/vaberof/smartway-task/internal/infra/storage"
	"github.com/vaberof/smartway-task/pkg/logging/logs"
	"log/slog"
)

var (
	ErrEmployeeNotFound   = errors.New("employee not found")
	ErrCompanyNotFound    = errors.New("company not found")
	ErrDepartmentNotFound = errors.New("department not found")
)

const errorEmployeeId = 0

type EmployeeService interface {
	Create(name, surname, phone string, companyId int64, passportType, passportNumber, departmentName, departmentPhone string) (int64, error)
	Update(id int64, name, surname, phone *string, companyId *int64, passportType, passportNumber, departmentName, departmentPhone *string) error
	Delete(id int64) error
	ListByCompanyId(companyId int64, limit, offset int) ([]*Employee, error)
	ListByDepartmentName(companyId int64, departmentName string, limit, offset int) ([]*Employee, error)
}

type employeeServiceImpl struct {
	employeeStorage EmployeeStorage

	logger *slog.Logger
}

func NewEmployeeService(employeeStorage EmployeeStorage, logs *logs.Logs) EmployeeService {
	logger := logs.WithName("domain.employee.service")
	return &employeeServiceImpl{
		employeeStorage: employeeStorage,
		logger:          logger,
	}
}

func (e *employeeServiceImpl) Create(name, surname, phone string, companyId int64, passportType, passportNumber, departmentName, departmentPhone string) (int64, error) {
	const operation = "Create"

	log := e.logger.With(
		slog.String("operation", operation),
		slog.String("name", name),
		slog.String("surname", surname),
		slog.Int64("companyId", companyId),
	)

	log.Info("creating an employee")

	employeeId, err := e.employeeStorage.Create(name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone)
	if err != nil {
		log.Error("failed to create an employee")

		if errors.Is(err, storage.ErrEmployeeNotFound) {
			return errorEmployeeId, fmt.Errorf("%s: %w", operation, ErrEmployeeNotFound)
		}
		return errorEmployeeId, fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("created an employee")

	return employeeId, nil
}

func (e *employeeServiceImpl) Update(id int64, name, surname, phone *string, companyId *int64, passportType, passportNumber, departmentName, departmentPhone *string) error {
	const operation = "Update"

	log := e.logger.With(
		slog.String("operation", operation),
		slog.Int64("id", id),
	)

	log.Info("updating an employee")

	err := e.employeeStorage.Update(id, name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone)
	if err != nil {
		log.Error("failed to update an employee")

		if errors.Is(err, storage.ErrEmployeeNotFound) {
			return fmt.Errorf("%s: %w", operation, ErrEmployeeNotFound)
		}
		return fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("updated an employee")

	return nil
}

func (e *employeeServiceImpl) Delete(id int64) error {
	const operation = "Delete"

	log := e.logger.With(
		slog.String("operation", operation),
		slog.Int64("id", id),
	)

	log.Info("deleting an employee")

	err := e.employeeStorage.Delete(id)
	if err != nil {
		log.Error("failed to delete an employee")

		if errors.Is(err, storage.ErrEmployeeNotFound) {
			return fmt.Errorf("%s: %w", operation, ErrEmployeeNotFound)
		}
		return fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("deleted an employee")

	return nil
}

func (e *employeeServiceImpl) ListByCompanyId(companyId int64, limit, offset int) ([]*Employee, error) {
	const operation = "ListByCompanyId"

	log := e.logger.With(
		slog.String("operation", operation),
		slog.Int64("companyId", companyId),
		slog.Int("limit", limit),
		slog.Int("offset", offset),
	)

	log.Info("listing employees by company id")

	employees, err := e.employeeStorage.ListByCompanyId(companyId, limit, offset)
	if err != nil {
		log.Error("failed to list employees by company id")

		if errors.Is(err, storage.ErrCompanyNotFound) {
			return nil, fmt.Errorf("%s: %w", operation, ErrCompanyNotFound)
		}
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("listed employees by company id")

	return employees, nil
}

func (e *employeeServiceImpl) ListByDepartmentName(companyId int64, departmentName string, limit, offset int) ([]*Employee, error) {
	const operation = "ListByDepartmentName"

	log := e.logger.With(
		slog.String("operation", operation),
		slog.Int64("companyId", companyId),
		slog.String("departmentName", departmentName),
		slog.Int("limit", limit),
		slog.Int("offset", offset),
	)

	log.Info("listing employees by department name")

	employees, err := e.employeeStorage.ListByDepartmentName(companyId, departmentName, limit, offset)
	if err != nil {
		log.Error("failed to list employees by department name")

		if errors.Is(err, storage.ErrDepartmentNotFound) {
			return nil, fmt.Errorf("%s: %w", operation, ErrDepartmentNotFound)
		}
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("listed employees by department name")

	return employees, nil
}
