CREATE TABLE tweet_stream (
    tweet_id BIGINT NOT NULL
    , stream_id BIGINT NOT NULL
    , PRIMARY KEY (tweet_id, stream_id)
);