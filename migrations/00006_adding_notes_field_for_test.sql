-- +goose Up
-- +goose StatementBegin

ALTER TABLE users
ADD COLUMN notes TEXT DEFAULT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin 
DROP TABLE users;
-- +goose StatementEnd