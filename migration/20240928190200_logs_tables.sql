-- +goose Up
-- +goose StatementBegin
CREATE TYPE log_level AS ENUM (
    'CRITICAL',
    'ERROR',
    'WARNING',
    'NOTICE',
    'INFO',
    'DEBUG'
);

CREATE TABLE IF NOT EXISTS modules (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS logs (
    id BIGSERIAL PRIMARY KEY,
    trace_id VARCHAR(255) NOT NULL,
    module_id BIGINT REFERENCES modules (id) NOT NULL,

    time TIMESTAMP NOT NULL,
    level log_level NOT NULL,
    message TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE logs;
DROP TABLE modules;
DROP TYPE log_level;
-- +goose StatementEnd
