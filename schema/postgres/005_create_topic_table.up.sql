CREATE TABLE topic (
    topic_id BIGSERIAL PRIMARY KEY
    , user_id BIGINT REFERENCES "user"(user_id) ON DELETE CASCADE
    , name TEXT NOT NULL
    , track TEXT NOT NULL
    , created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE stream
    ADD COLUMN topic_id BIGINT;

ALTER TABLE stream
    ADD CONSTRAINT stream_topic_fk FOREIGN KEY (topic_id) REFERENCES topic (topic_id) ON DELETE CASCADE;