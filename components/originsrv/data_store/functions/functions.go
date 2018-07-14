package functions

// This whole package is about pseudo stored procedures since cockroach can't handle them yet.

const GetOriginByNameV1 = "SELECT * FROM origins WHERE name = $1"
const InsertOriginV1 = "INSERT INTO origins(name, default_package_visibility) VALUES ($1, $2) RETURNING id"
const UpdateOriginV1 = "UPDATE origins SET default_package_visibility = $1 WHERE name = $2 RETURNING *"

const GetProjectV1 = "SELECT * FROM projects WHERE origin = $1 AND package_name = $2"
const InsertProjectV1 = "INSERT INTO projects (origin, package_name, plan_path, vcs_type, vcs_data) VALUES ($1, $2, $3, $4, $5) RETURNING *"
const UpdateProjectV1 = "UPDATE projects SET plan_path = $3, vcs_type = $4, vcs_data = $5 WHERE origin = $1 AND package_name = $2 RETURNING *"
