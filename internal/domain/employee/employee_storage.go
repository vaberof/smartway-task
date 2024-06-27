package employee

type EmployeeStorage interface {
	Create(name, surname, phone string, companyId int64, passportType, passportNumber string, departmentName, departmentPhone string) (int64, error)
	Update(id int64, name, surname, phone *string, companyId *int64, passportType, passportNumber, departmentName, departmentPhone *string) error
	Delete(id int64) error
	ListByCompanyId(companyId int64, limit, offset int) ([]*Employee, error)
	ListByDepartmentName(companyId int64, departmentName string, limit, offset int) ([]*Employee, error)
}
