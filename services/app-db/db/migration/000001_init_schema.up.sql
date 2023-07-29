-- name: CreateUsers :one

CREATE TABLE "User" (
                        "Username" varchar UNIQUE,
                        "Full_Name" varchar,
                        "Email" varchar UNIQUE,
                        "Password_Changed_At" timestamp,
                        "Created_At" timestamp
);

CREATE TABLE "ExpenseType" (
                               "ExpenseTypeID" int PRIMARY KEY,
                               "Name" varchar UNIQUE,
                               "Description" varchar
);

CREATE TABLE "Expense" (
                           "ExpenseID" int PRIMARY KEY,
                           "ExpenseTypeID" int,
                           "Amount" int,
                           "Description" varchar,
                           "Created_At" timestamp
);

ALTER TABLE "Expense" ADD FOREIGN KEY ("ExpenseTypeID") REFERENCES "ExpenseType" ("ExpenseTypeID");
