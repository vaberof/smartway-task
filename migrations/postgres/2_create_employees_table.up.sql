CREATE TABLE IF NOT EXISTS employees
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(30) NOT NULL,
    surname    VARCHAR(30) NOT NULL,
    phone      VARCHAR(15) NOT NULL,
    company_id INT         NOT NULL REFERENCES companies (id) ON DELETE CASCADE,
    passport   JSONB,
    department JSONB
);
CREATE INDEX IF NOT EXISTS company_id_idx ON employees (company_id);
CREATE INDEX IF NOT EXISTS department_name_idx ON employees ((department ->> 'name'));
