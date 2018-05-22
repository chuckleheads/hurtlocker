package functions

// This whole package is about pseudo stored procedures since cockroach can't handle them yet.

const InsertAccountV1 = "INSERT INTO accounts(username, email) VALUES($1, $2) RETURNING id"
const GetAccountByUsernameV1 = "SELECT * FROM accounts WHERE username = $1"
const UpdateAccountV1 = "UPDATE accounts SET email = $1 WHERE username = $2 RETURNING *"
