-- name: GetExpenseTypeByID :one
SELECT * FROM "ExpenseType"
WHERE "ExpenseTypeID" = $1
    LIMIT 1;

-- name: ListExpenseTypes :many
SELECT * FROM "ExpenseType"
ORDER BY "Name";

-- name: CreateExpenseType :one
INSERT INTO "ExpenseType" (
    "ExpenseTypeID", "Name", "Description"
) VALUES (
             $1, $2, $3
         )
    RETURNING *;

-- name: UpdateExpenseType :one
UPDATE "ExpenseType"
SET "Name" = $2, "Description" = $3
WHERE "ExpenseTypeID" = $1
    RETURNING *;

-- name: DeleteExpenseType :exec
DELETE FROM "ExpenseType"
WHERE "ExpenseTypeID" = $1;
