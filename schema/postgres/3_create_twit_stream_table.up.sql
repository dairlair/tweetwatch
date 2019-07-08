CREATE TABLE twit_stream (
    twit_id BIGINT NOT NULL
    , stream_id BIGINT NOT NULL
    , PRIMARY KEY (twit_id, stream_id)
);