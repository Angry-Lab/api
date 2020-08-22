CREATE TABLE IF NOT EXISTS city_index
(
    id        SERIAL            NOT NULL,
    postcode  INT               NOT NULL,
    name      CHARACTER VARYING NOT NULL,
    latitude  NUMERIC           NOT NULL,
    longitude NUMERIC           NOT NULL,
    CONSTRAINT city_index_pkey PRIMARY KEY (id),
    CONSTRAINT city_index_ukey UNIQUE (postcode)
);


