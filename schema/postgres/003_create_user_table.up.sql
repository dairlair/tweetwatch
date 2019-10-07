CREATE TABLE "user" (
    user_id BIGSERIAL PRIMARY KEY
    , email TEXT
    , hash TEXT
    , created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_email_unique_idx on "user" (LOWER(email));