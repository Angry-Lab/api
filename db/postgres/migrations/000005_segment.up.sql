CREATE TABLE IF NOT EXISTS segment (
    id SERIAL NOT NULL,
    name CHARACTER VARYING NOT NULL,
    description CHARACTER VARYING NOT NULL,
    condition CHARACTER VARYING NOT NULL,
    CONSTRAINT segment_pkey PRIMARY KEY (id)
);
