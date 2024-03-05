CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    role_id INTEGER,
    phonenumber VARCHAR(12),
    email VARCHAR(255),
    password VARCHAR(255),
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE 
);