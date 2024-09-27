-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS logs (
    id BIGSERIAL PRIMARY KEY,
    trace_id VARCHAR(255) NOT NULL,

    time TIMESTAMP NOT NULL,
    module VARCHAR(255) NOT NULL,
    level INT NOT NULL,

    message TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE logs;
-- +goose StatementEnd
