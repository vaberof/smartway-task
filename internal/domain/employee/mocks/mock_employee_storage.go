// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/employee/employee_storage.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/employee/employee_storage.go -destination=internal/domain/employee/mocks/mock_employee_storage.go
//

// Package mock_employee is a generated GoMock package.
package mock_employee

import (
	reflect "reflect"

	employee "github.com/vaberof/smartway-task/internal/domain/employee"
	gomock "go.uber.org/mock/gomock"
)

// MockEmployeeStorage is a mock of EmployeeStorage interface.
type MockEmployeeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeStorageMockRecorder
}

// MockEmployeeStorageMockRecorder is the mock recorder for MockEmployeeStorage.
type MockEmployeeStorageMockRecorder struct {
	mock *MockEmployeeStorage
}

// NewMockEmployeeStorage creates a new mock instance.
func NewMockEmployeeStorage(ctrl *gomock.Controller) *MockEmployeeStorage {
	mock := &MockEmployeeStorage{ctrl: ctrl}
	mock.recorder = &MockEmployeeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeStorage) EXPECT() *MockEmployeeStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEmployeeStorage) Create(name, surname, phone string, companyId int64, passportType, passportNumber, departmentName, departmentPhone string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEmployeeStorageMockRecorder) Create(name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEmployeeStorage)(nil).Create), name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone)
}

// Delete mocks base method.
func (m *MockEmployeeStorage) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEmployeeStorageMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEmployeeStorage)(nil).Delete), id)
}

// ListByCompanyId mocks base method.
func (m *MockEmployeeStorage) ListByCompanyId(companyId int64, limit, offset int) ([]*employee.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByCompanyId", companyId, limit, offset)
	ret0, _ := ret[0].([]*employee.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByCompanyId indicates an expected call of ListByCompanyId.
func (mr *MockEmployeeStorageMockRecorder) ListByCompanyId(companyId, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByCompanyId", reflect.TypeOf((*MockEmployeeStorage)(nil).ListByCompanyId), companyId, limit, offset)
}

// ListByDepartmentName mocks base method.
func (m *MockEmployeeStorage) ListByDepartmentName(companyId int64, departmentName string, limit, offset int) ([]*employee.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDepartmentName", companyId, departmentName, limit, offset)
	ret0, _ := ret[0].([]*employee.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByDepartmentName indicates an expected call of ListByDepartmentName.
func (mr *MockEmployeeStorageMockRecorder) ListByDepartmentName(companyId, departmentName, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDepartmentName", reflect.TypeOf((*MockEmployeeStorage)(nil).ListByDepartmentName), companyId, departmentName, limit, offset)
}

// Update mocks base method.
func (m *MockEmployeeStorage) Update(id int64, name, surname, phone *string, companyId *int64, passportType, passportNumber, departmentName, departmentPhone *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeStorageMockRecorder) Update(id, name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeStorage)(nil).Update), id, name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone)
}
