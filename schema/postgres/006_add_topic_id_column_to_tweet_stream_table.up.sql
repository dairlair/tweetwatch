TRUNCATE TABLE tweet_stream;

ALTER TABLE tweet_stream
    ADD COLUMN topic_id BIGINT NOT NULL,
    ADD CONSTRAINT tweet_stream_topic_fk FOREIGN KEY (topic_id) REFERENCES topic (topic_id) ON DELETE CASCADE;