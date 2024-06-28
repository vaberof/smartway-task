package employee

type CompanyStorage interface {
	IsExists(id int64) (bool, error)
}
