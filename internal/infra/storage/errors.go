package storage

import "errors"

var (
	ErrEmployeeNotFound   = errors.New("employee not found")
	ErrCompanyNotFound    = errors.New("company not found")
	ErrDepartmentNotFound = errors.New("department not found")
)
