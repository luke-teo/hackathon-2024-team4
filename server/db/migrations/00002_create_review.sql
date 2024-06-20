-- +goose Up
-- +goose StatementBegin
CREATE TABLE "review" (
  "id" bigserial PRIMARY KEY,
  "rating" int,
  "content" text,
  "user_id" bigint NOT NULL,
  "tenant_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE INDEX ON "review" ("rating");

ALTER TABLE "review" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
ALTER TABLE "review" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "review";
-- +goose StatementEnd
