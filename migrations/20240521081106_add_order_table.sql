-- Create "orders" table
CREATE TABLE "orders" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "drone_serial_number" bigint NULL,
  "medication_code" text NULL,
  "quantity" bigint NULL,
  "state" text NOT NULL DEFAULT 'PROCESSING',
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_orders_drone" FOREIGN KEY ("drone_serial_number") REFERENCES "drones" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_orders_medication" FOREIGN KEY ("medication_code") REFERENCES "medications" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_orders_deleted_at" to table: "orders"
CREATE INDEX "idx_orders_deleted_at" ON "orders" ("deleted_at");
