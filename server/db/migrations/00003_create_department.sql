-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION ltree;

CREATE TABLE "department" (
  "id" bigserial PRIMARY KEY,
  "name" text,
  "custom_id" text NOT NULL,
  "tenant_id" int NOT NULL,
  "hierarchy" ltree NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

ALTER TABLE "department" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "department";

DROP EXTENSION ltree;
-- +goose StatementEnd
