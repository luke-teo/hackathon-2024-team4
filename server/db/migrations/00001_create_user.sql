-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "name" text,
  "email" text NOT NULL,
  "role" int NOT NULL,
  "tenant_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "tenant" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "short_code" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE UNIQUE INDEX ON "user" ("email", "tenant_id");

CREATE INDEX ON "user" ("role");

CREATE UNIQUE INDEX ON "tenant" ("short_code");

ALTER TABLE "user" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user";

DROP TABLE "tenant";
-- +goose StatementEnd
