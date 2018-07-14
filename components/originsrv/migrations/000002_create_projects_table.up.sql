CREATE TABLE IF NOT EXISTS projects (
    id bigserial PRIMARY KEY,
    origin TEXT NOT NULL,
		package_name TEXT NOT NULL,
		plan_path TEXT NOT NULL,
		vcs_type TEXT NOT NULL,
		vcs_data TEXT NOT NULL,
		UNIQUE(origin, package_name)
)
