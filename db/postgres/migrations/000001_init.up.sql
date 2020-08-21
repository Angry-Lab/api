CREATE TYPE USER_STATUS AS ENUM ('new', 'active', 'blocked');
CREATE TYPE USER_ROLE AS ENUM ('owner', 'root', 'user');

CREATE TABLE IF NOT EXISTS "user"
(
    id             BIGSERIAL                   NOT NULL,
    login          CHARACTER VARYING(128)      NOT NULL,
    password       CHARACTER(64)               NOT NULL,
    status         USER_STATUS                 NOT NULL DEFAULT 'new',
    name           CHARACTER VARYING(255)      NOT NULL,
    role           USER_ROLE                   NOT NULL,
    dt_created     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    dt_updated     TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    dt_last_logged TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_login_ukey UNIQUE (login)
);

CREATE TYPE TOKEN_TYPE AS ENUM ('auth');

CREATE TABLE IF NOT EXISTS auth_token
(
    id         BIGSERIAL                   NOT NULL,
    user_id    BIGINT                      NOT NULL,
    token      CHARACTER(64)               NOT NULL,
    type       TOKEN_TYPE                  NOT NULL,
    dt_expired TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    dt_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT auth_token_pkey PRIMARY KEY (id),
    CONSTRAINT auth_token_ukey UNIQUE (token),
    CONSTRAINT user_auth_token_fkey FOREIGN KEY (user_id) REFERENCES "user" (id)
        ON DELETE CASCADE
);

-- INSERT INTO "user" ("id", "login", "password", "status", "name", "role",
--                     "dt_created", "dt_updated", "dt_last_logged")
-- VALUES (1, 'partyzan65@gmail.com', '$2a$04$upglOQmqyjs.44Hj.YLrf.8NKpkAi4Kr.v92mLhKV9/vuYmeHv5nq    ', 'active',
--         'Admin', 'owner', now(), now(), '0001-01-01 00:00:00');
