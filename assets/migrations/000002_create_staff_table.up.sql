CREATE TABLE IF NOT EXISTS staff (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL,
    start_date date NOT NULL,
    team_id bigint NOT NULL REFERENCES teams,
    version integer NOT NULL DEFAULT 1
);