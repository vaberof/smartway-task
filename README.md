## Использование

    make docker.run

## Swagger

    http://localhost:8000/swagger/index.html

## Юнит-тесты

    make tests.run

    make tests.run.verbose

## Решения по схеме базы данных

1. Можно было создать таблицы: companies(id, name), departments(id, name, phone, company_id), employees(id, name,
   surname, phone, company_id, department_id, passport_id), employees_passports(number, type, employee_id)
2. Или: companies(id, name), employees(id, name, surname, phone, company_id), employees_details(employee_id,
   passport_type, passport_number, department_name, department_phone)

Но в итоге остановился на текущей версии: companies(id, name), employees(id, name, surname, phone, company_id, passport [JSONB],
department [JSONB]), так как не понял из описанной модели в ТЗ, можно ли передавать department_id в запросе. Если нельзя - то с текущей версией проще было реализовать нужную логику.
Если бы можно было передавать department_id, то выбрал бы первый вариант, так как он кажется наиболее правильным с точки зрения структуры
