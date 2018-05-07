CREATE TABLE IF NOT EXISTS origins (
    id bigserial PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    default_package_visibility text
    NOT NULL DEFAULT 'public'
)
