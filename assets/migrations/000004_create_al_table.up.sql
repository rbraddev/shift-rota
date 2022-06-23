CREATE TABLE IF NOT EXISTS leave (
    id bigserial PRIMARY KEY,
    start_date timestamp NOT NULL,
    end_date timestamp NOT NULL,
    toil bool NOT NULL DEFAULT false,
    notes text,
    staff_id bigint NOT NULL REFERENCES staff,
    version integer NOT NULL DEFAULT 1
);