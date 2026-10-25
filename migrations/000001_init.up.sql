CREATE SCHEMA IF NOT EXISTS todoapp;

CREATE TABLE todoapp.users(
    id           SERIAL         PRIMARY KEY,
    version      BIGINT         NOT NULL DEFAULT 1,
    full_name    VARCHAR(100)   NOT NULL CHECK (char_length(full_name) between 3 and 100),
    phone_number VARCHAR(15)    CHECK(phone_number ~ '^\+[0-9]+$' AND char_length(phone_number) between 10 and 15)
);

CREATE TABLE todoapp.tasks(
    id              SERIAL          PRIMARY KEY,
    version         BIGINT          NOT NULL DEFAULT 1,
    title           VARCHAR(100)    NOT NULL check (char_length(title) between 1 and 100),
    description     VARCHAR(1000)   check ( char_length(description) between 1 and 1000),
    completed       BOOLEAN         NOT NULL,
    created_at      TIMESTAMPTZ     NOT NULL,
    completed_at    TIMESTAMPTZ,

    check (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    ),

    author_user_id INTEGER NOT NULL REFERENCES todoapp.users(id)
);


