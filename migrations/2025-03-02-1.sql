BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS vacancies (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    email VARCHAR(255) NOT NULL,
    salary VARCHAR(255),
    company VARCHAR(255),
    role VARCHAR(200),
    location VARCHAR(255),
    direction VARCHAR(255)
)

COMMIT;