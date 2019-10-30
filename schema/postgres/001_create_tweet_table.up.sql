CREATE TABLE tweet (
    tweet_id BIGSERIAL PRIMARY KEY
    , twitter_id BIGINT NOT NULL
    , twitter_user_id BIGINT NOT NULL
    , twitter_username TEXT NOT NULL
    , full_text TEXT NOT NULL
    , created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);