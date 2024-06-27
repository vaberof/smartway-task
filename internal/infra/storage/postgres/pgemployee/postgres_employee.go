package pgemployee

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type PgEmployee struct {
	Id         int64
	Name       string
	Surname    string
	Phone      string
	CompanyId  int64
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type Passport struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

func (p *Passport) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Passport) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)
}

type Department struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (d *Department) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *Department) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &d)
}
