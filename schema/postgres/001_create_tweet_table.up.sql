CREATE TABLE tweet (
    tweet_id BIGSERIAL PRIMARY KEY
    , id BIGINT NOT NULL
    , user_id BIGINT NOT NULL
    , full_text TEXT NOT NULL
    , created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN tweet.id IS 'The tweet identifier from Twitter database';