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


