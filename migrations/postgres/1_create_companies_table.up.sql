CREATE TABLE IF NOT EXISTS companies
(
    id   SERIAL PRIMARY KEY,
    name TEXT
);

INSERT INTO companies (name)
VALUES ('Company_1');
INSERT INTO companies (name)
VALUES ('Company_2');
INSERT INTO companies (name)
VALUES ('Company_3');