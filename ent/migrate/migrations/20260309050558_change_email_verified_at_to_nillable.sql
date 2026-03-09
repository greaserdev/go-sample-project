-- Modify "credentials" table
ALTER TABLE "public"."credentials" ALTER COLUMN "email_verified_at" DROP NOT NULL;
-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "verified_at" DROP NOT NULL, ALTER COLUMN "avatar_url" DROP NOT NULL;
