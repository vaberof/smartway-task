package employee_service

import (
	"github.com/stretchr/testify/require"
	"github.com/vaberof/smartway-task/internal/domain/employee"
	mocks "github.com/vaberof/smartway-task/internal/domain/employee/mocks"
	"github.com/vaberof/smartway-task/pkg/logging/logs"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeStorage := mocks.NewMockEmployeeStorage(ctrl)
	logsBuilder := logs.New(os.Stdout, nil)

	employeeService := employee.NewEmployeeService(employeeStorage, logsBuilder)

	name := "name"
	surname := "surname"
	phone := "phone"
	companyId := int64(1)
	passportType := "passportType"
	passportNumber := "passportNumber"
	departmentName := "departmentName"
	departmentPhone := "departmentPhone"

	expectedId := int64(1)

	employeeStorage.EXPECT().Create(name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone).Return(expectedId, nil).AnyTimes()
	actualId, err := employeeService.Create(name, surname, phone, companyId, passportType, passportNumber, departmentName, departmentPhone)

	require.NoError(t, err)
	require.Equal(t, expectedId, actualId)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeStorage := mocks.NewMockEmployeeStorage(ctrl)
	logsBuilder := logs.New(os.Stdout, nil)

	employeeService := employee.NewEmployeeService(employeeStorage, logsBuilder)

	employeeId := int64(1)

	employeeStorage.EXPECT().Delete(employeeId).Return(nil).AnyTimes()
	actualError := employeeService.Delete(employeeId)

	require.NoError(t, actualError)
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeStorage := mocks.NewMockEmployeeStorage(ctrl)
	logsBuilder := logs.New(os.Stdout, nil)

	employeeService := employee.NewEmployeeService(employeeStorage, logsBuilder)

	id := int64(1)
	name := "name"
	surname := "surname"
	phone := "phone"
	companyId := int64(1)
	passportType := "passportType"
	passportNumber := "passportNumber"
	departmentName := "departmentName"
	departmentPhone := "departmentPhone"

	employeeStorage.EXPECT().Update(id, &name, &surname, &phone, &companyId, &passportType, &passportNumber, &departmentName, &departmentPhone).Return(nil).AnyTimes()
	err := employeeService.Update(id, &name, &surname, &phone, &companyId, &passportType, &passportNumber, &departmentName, &departmentPhone)

	require.NoError(t, err)
}

func TestListByCompanyId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeStorage := mocks.NewMockEmployeeStorage(ctrl)
	logsBuilder := logs.New(os.Stdout, nil)

	employeeService := employee.NewEmployeeService(employeeStorage, logsBuilder)

	expectedEmployees := []*employee.Employee{
		{
			Id:        1,
			Name:      "name_1",
			Surname:   "surname_1",
			Phone:     "phone_1",
			CompanyId: int64(1),
			Passport: employee.Passport{
				Type:   "passportType_1",
				Number: "passportNumber_1",
			},
			Department: employee.Department{
				Name:  "departmentName_1",
				Phone: "departmentPhone_1",
			},
		},
		{
			Id:        2,
			Name:      "name_2",
			Surname:   "surname_2",
			Phone:     "phone_2",
			CompanyId: int64(1),
			Passport: employee.Passport{
				Type:   "passportType_2",
				Number: "passportNumber_2",
			},
			Department: employee.Department{
				Name:  "departmentName_2",
				Phone: "departmentPhone_2",
			},
		},
	}

	listByCompanyId := int64(1)
	limit := 10
	offset := 0

	employeeStorage.EXPECT().ListByCompanyId(listByCompanyId, limit, offset).Return(expectedEmployees, nil).AnyTimes()
	actualEmployees, err := employeeService.ListByCompanyId(listByCompanyId, limit, offset)

	require.NoError(t, err)
	require.Equal(t, expectedEmployees, actualEmployees)
}

func TestListByDepartmentName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeStorage := mocks.NewMockEmployeeStorage(ctrl)
	logsBuilder := logs.New(os.Stdout, nil)

	employeeService := employee.NewEmployeeService(employeeStorage, logsBuilder)

	expectedEmployees := []*employee.Employee{
		{
			Id:        1,
			Name:      "name_1",
			Surname:   "surname_1",
			Phone:     "phone_1",
			CompanyId: int64(1),
			Passport: employee.Passport{
				Type:   "passportType_1",
				Number: "passportNumber_1",
			},
			Department: employee.Department{
				Name:  "departmentName_1",
				Phone: "departmentPhone_1",
			},
		},
		{
			Id:        2,
			Name:      "name_2",
			Surname:   "surname_2",
			Phone:     "phone_2",
			CompanyId: int64(1),
			Passport: employee.Passport{
				Type:   "passportType_2",
				Number: "passportNumber_2",
			},
			Department: employee.Department{
				Name:  "departmentName_1",
				Phone: "departmentPhone_1",
			},
		},
	}

	listByCompanyId := int64(1)
	listByDepartmentName := "departmentName_1"
	limit := 10
	offset := 0

	employeeStorage.EXPECT().ListByDepartmentName(listByCompanyId, listByDepartmentName, limit, offset).Return(expectedEmployees, nil).AnyTimes()
	actualEmployees, err := employeeService.ListByDepartmentName(listByCompanyId, listByDepartmentName, limit, offset)

	require.NoError(t, err)
	require.Equal(t, expectedEmployees, actualEmployees)
}
