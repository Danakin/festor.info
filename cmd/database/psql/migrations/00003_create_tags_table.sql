-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags (
    "id" SERIAL PRIMARY KEY,
    "title" varchar NOT NULL,
    "font_color" varchar(7),
    "background_color" varchar(7),
    "border_color" varchar(7),
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now()
);
CREATE TABLE post_tag (
    "post_id" int8,
    "tag_id" int8,
    CONSTRAINT "post_tag_post_id_fkey" FOREIGN KEY ("post_id") REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT "post_tag_tag_id_fkey" FOREIGN KEY ("tag_id") REFERENCES tags(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_tag;
DROP TABLE IF EXISTS tags;
-- +goose StatementEnd