-- name: CreateRole :exec
INSERT INTO roles (name)
VALUES ($1);

-- name: GetRoles :many
SELECT id, name
FROM roles;

-- name: GetRoleByID :one
SELECT id, name
FROM roles
WHERE id = $1;

-- name: GetRoleByName :one
SELECT id, name
FROM roles
WHERE name = $1;

-- name: UpdateRole :exec
UPDATE roles
SET name = $1
WHERE id = $2;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;

-- name: CreateUserRole :exec
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2);

-- name: GetUserRoleByID :one
SELECT id, user_id, role_id
FROM user_roles
WHERE id = $1;

-- name: GetUserRoleByUserID :one
SELECT 
    ur.id AS user_role_id,
    ur.user_id,
    ur.role_id,
    r.id AS role_id,
    r.name AS role_name
FROM user_roles ur
LEFT JOIN roles r ON ur.role_id = r.id
WHERE ur.user_id = $1;

-- name: UpdateUserRole :exec
UPDATE user_roles
SET user_id = $1, role_id = $2
WHERE id = $3;

-- name: DeleteUserRole :exec
DELETE FROM user_roles
WHERE id = $1;

