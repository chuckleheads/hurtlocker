package functions

// This whole package is about pseudo stored procedures since cockroach can't handle them yet.

const GetJobByIdV1 = "SELECT * FROM jobs WHERE id = $1"
const InsertJobV1 = "INSERT INTO jobs(name, default_package_visibility) VALUES ($1, $2) RETURNING id" // wrong
