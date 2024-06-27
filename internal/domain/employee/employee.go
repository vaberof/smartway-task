package employee

type Employee struct {
	Id         int64
	Name       string
	Surname    string
	Phone      string
	CompanyId  int64
	Passport   Passport
	Department Department
}

type Passport struct {
	Type   string
	Number string
}

type Department struct {
	Name  string
	Phone string
}
