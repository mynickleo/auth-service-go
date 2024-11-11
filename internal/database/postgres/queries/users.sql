-- name: CreateUser :exec
INSERT INTO users (id, full_name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUsers :many
SELECT id, email, full_name, created_at, updated_at, avatar_img
FROM users;

-- name: GetUserByID :one
SELECT id, email, full_name, created_at, updated_at, avatar_img
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT u.id AS user_id, u.email, u.password, u.full_name, u.created_at, u.updated_at, avatar_img, r.name AS role_name
FROM users u
LEFT JOIN user_roles ur ON ur.user_id = u.id
LEFT JOIN roles r ON r.id = ur.role_id
WHERE u.email = $1;

-- name: UpdateUser :exec
UPDATE users
SET email = $1, password = $2, full_name = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4;

-- name: UpdateUserAvatar :exec
UPDATE users
SET avatar_img = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
