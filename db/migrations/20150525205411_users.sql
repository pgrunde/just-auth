-- Auth user with sessions

-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "users" (
  "id" SERIAL NOT NULL,
  "email" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "about" VARCHAR,
  "is_active" BOOLEAN NOT NULL DEFAULT TRUE,
  "is_superuser" BOOLEAN NOT NULL DEFAULT FALSE,
  "password" VARCHAR NOT NULL,
  "token" VARCHAR NOT NULL,
  "token_set_at" TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
  "created_at" TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
  "deleted_at" TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
  PRIMARY KEY ("id")
);
CREATE TABLE "sessions" (
  "key" VARCHAR NOT NULL,
  "user_id" INTEGER NOT NULL REFERENCES users("id") ON DELETE CASCADE,
  "expires" TIMESTAMP NOT NULL,
  PRIMARY KEY ("key")
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "sessions";
