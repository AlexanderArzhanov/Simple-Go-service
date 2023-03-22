CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    msg VARCHAR(64) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
    );

INSERT INTO events (msg, created_at)
VALUES ('first event', '2023-02-01 16:10:25-07');

INSERT INTO events (msg, created_at)
VALUES ('second event', '2023-02-03 16:10:25-07');

INSERT INTO events (msg, created_at)
VALUES ('third event', '2023-02-05 16:10:25-07');