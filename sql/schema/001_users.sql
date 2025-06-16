-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE staff (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL UNIQUE,
  role varchar(255) CHECK(role IN ('Doctor', 'Receptionist')) NOT NULL, 
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  pw_hash varchar(255) not null
);


-- +goose Down
DROP TABLE staff;
