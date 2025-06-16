-- name: CreateStaffMember :one
INSERT INTO staff (id, name, role, created_at, updated_at, pw_hash)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetStaffMember :one
SELECT * FROM staff 
WHERE name = $1;

-- name: DropRows :exec
TRUNCATE TABLE staff RESTART IDENTITY CASCADE;

-- name: ListStaffMembers :many
SELECT name FROM staff;

-- name: GetStaffPasswdHash :one
SELECT pw_hash from staff where name = $1;

-- name: CreatePatient :one
insert into patients (id, name, age, gender, diagnosis, address, created_at, updated_at)
values($1, $2, $3, $4, $5, $6, $7, $8)
Returning *;

-- name: GetPatient :one
Select * from patients
where name = $1;

-- name: DeletePatient :exec
DELETE from patients
where name = $1;

-- name: UpdatePatientDetails :one
UPDATE patients
SET name = $2, age = $3, gender = $4, address = $5, diagnosis = $6, updated_at = NOW()
WHERE name = $1
RETURNING *;
