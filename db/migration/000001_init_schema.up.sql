CREATE TABLE IF NOT EXISTS "candidates" (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    middle_name VARCHAR(50),
    date_of_birth DATE,
    city VARCHAR(50),
    education VARCHAR(255),
    reason_dismissal TEXT,
    email VARCHAR(255),
    phone VARCHAR(50),
    year_work_experience INTEGER,
    employee_entered_info VARCHAR(255)
);