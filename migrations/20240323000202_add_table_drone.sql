-- Create "drones" table
CREATE TABLE "drones" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "serial_number" character varying(100) NOT NULL,
  "drone_model" text NOT NULL,
  "weight_limit" numeric NOT NULL,
  "battery_capacity" bigint NOT NULL,
  "state" text NOT NULL DEFAULT 'IDLE',
  PRIMARY KEY ("id"),
  CONSTRAINT "chk_drones_battery_capacity" CHECK (battery_capacity <= 100),
  CONSTRAINT "chk_drones_weight_limit" CHECK (weight_limit <= (500)::numeric)
);
-- Create index "idx_drones_deleted_at" to table: "drones"
CREATE INDEX "idx_drones_deleted_at" ON "drones" ("deleted_at");
-- Create index "uni_drones_serial_number" to table: "drones"
CREATE UNIQUE INDEX "uni_drones_serial_number" ON "drones" ("serial_number");
