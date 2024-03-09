-- Create "drones" table
CREATE TABLE "drones" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "serial_number" character varying(100) NULL,
  "drone_model" text NULL,
  "weight_limit" bigint NULL,
  "battery_capacity" bigint NULL,
  "state" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "chk_drones_battery_capacity" CHECK (battery_capacity <= 100),
  CONSTRAINT "chk_drones_weight_limit" CHECK (weight_limit <= 500)
);
-- Create index "idx_drones_deleted_at" to table: "drones"
CREATE INDEX "idx_drones_deleted_at" ON "drones" ("deleted_at");
