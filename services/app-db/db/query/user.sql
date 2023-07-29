-- name: CreateUser :one
INSERT INTO "User" (
    "Username",
    "Full_Name",
    "Email",
    "Password_Changed_At",
    "Created_At"
) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING *;

-- name: GetUser :one
SELECT * FROM "User"
WHERE "Username" = $1 LIMIT 1;


-- name: UpdateUser :one
-- UPDATE User
-- SET
--     hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
--     password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
--     full_name = COALESCE(sqlc.narg(full_name), full_name),
--     email = COALESCE(sqlc.narg(email), email),
--     is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
-- WHERE
--         username = sqlc.arg(username)
-- RETURNING *;
