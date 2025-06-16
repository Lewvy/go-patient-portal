-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE patients (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL UNIQUE,
  age int NOT NULL,
  gender char(2) CHECK(gender IN ('M', 'F', 'NB', 'O')) NOT NULL,
  address varchar(255),
  diagnosis varchar(255),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);


-- +goose Down
DROP TABLE patients;
