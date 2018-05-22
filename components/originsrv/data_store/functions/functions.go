package functions

// This whole package is about pseudo stored procedures since cockroach can't handle them yet.

const GetOriginByNameV1 = "SELECT * FROM origins WHERE name = $1"
const InsertOriginV1 = "INSERT INTO origins(name, default_package_visibility) VALUES ($1, $2) RETURNING id"
const UpdateOriginV1 = "UPDATE origins SET default_package_visibility = $1 WHERE name = $2 RETURNING *"
