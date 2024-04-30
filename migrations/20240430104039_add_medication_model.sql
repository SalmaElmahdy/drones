-- Create "medications" table
CREATE TABLE "medications" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_medications_deleted_at" to table: "medications"
CREATE INDEX "idx_medications_deleted_at" ON "medications" ("deleted_at");
