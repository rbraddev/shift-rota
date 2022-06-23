CREATE TABLE IF NOT EXISTS rotas (
    id bigserial PRIMARY KEY,
    rota json NOT NULL,
    start_date date NOT NULL,
    staff_id bigint NOT NULL REFERENCES staff,
    version integer NOT NULL DEFAULT 1
);