-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_behavior" (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  user_id TEXT NOT NULL,
  date DATE NOT NULL,
  score INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX "user_behavior_user_id_idx" ON "user_behavior" ("user_id");
CREATE INDEX "user_behavior_date_idx" ON "user_behavior" ("date");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "user_behavior_user_id_idx";
DROP INDEX "user_behavior_date_idx";

DROP TABLE "user_behavior";
-- +goose StatementEnd
