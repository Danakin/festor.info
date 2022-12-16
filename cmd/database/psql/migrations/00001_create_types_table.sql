-- +goose Up
-- +goose StatementBegin
CREATE TABLE types (
    "id" SERIAL PRIMARY KEY,
    "title" varchar NOT NULL,
    "description" text NOT NULL,
    "logo" varchar,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS types;
-- +goose StatementEnd
