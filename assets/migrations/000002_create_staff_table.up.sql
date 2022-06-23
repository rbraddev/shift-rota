CREATE TABLE IF NOT EXISTS staff (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    start_date date NOT NULL,
    team_id bigint NOT NULL REFERENCES teams,
    version integer NOT NULL DEFAULT 1
);