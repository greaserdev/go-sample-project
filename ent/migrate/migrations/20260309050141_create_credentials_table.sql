-- Create "credentials" table
CREATE TABLE "public"."credentials" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  "email" character varying NOT NULL,
  "password_hash" character varying NOT NULL,
  "email_verified_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "credentials_user" uuid NULL, ADD CONSTRAINT "users_credentials_user" FOREIGN KEY ("credentials_user") REFERENCES "public"."credentials" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Create index "users_credentials_user_key" to table: "users"
CREATE UNIQUE INDEX "users_credentials_user_key" ON "public"."users" ("credentials_user");
