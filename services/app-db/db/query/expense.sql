-- name: GetExpenseByID :one
SELECT * FROM "Expense"
WHERE "ExpenseID" = $1
    LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM "Expense"
ORDER BY "Created_At" DESC;

-- name: CreateExpense :one
INSERT INTO "Expense" (
    "ExpenseID", "ExpenseTypeID", "Amount", "Description", "Created_At"
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: UpdateExpense :one
UPDATE "Expense"
SET "ExpenseTypeID" = $2, "Amount" = $3, "Description" = $4, "Created_At" = $5
WHERE "ExpenseID" = $1
    RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM "Expense"
WHERE "ExpenseID" = $1;
