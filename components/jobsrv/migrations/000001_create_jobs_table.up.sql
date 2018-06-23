CREATE TABLE IF NOT EXISTS jobs (
  id bigserial PRIMARY KEY,
  owner_id bigint,
  project_id bigint,
  job_type text
)
