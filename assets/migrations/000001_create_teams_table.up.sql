CREATE TABLE IF NOT EXISTS teams (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    version integer NOT NULL DEFAULT 1
);