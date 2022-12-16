-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    "id" SERIAL PRIMARY KEY,
    "type_id" int8,
    "title" varchar NOT NULL,
    "description" text NOT NULL,
    "image" varchar,
    "is_released" bool NOT NULL DEFAULT false,
    "released_at" timestamp,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    CONSTRAINT "posts_type_id_fkey" FOREIGN KEY ("type_id") REFERENCES types(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd