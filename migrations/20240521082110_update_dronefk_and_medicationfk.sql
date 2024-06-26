-- Modify "orders" table
ALTER TABLE "orders" DROP CONSTRAINT "fk_orders_drone", DROP CONSTRAINT "fk_orders_medication", ADD CONSTRAINT "chk_orders_quantity" CHECK (quantity >= 0), DROP COLUMN "drone_serial_number", DROP COLUMN "medication_code", ADD COLUMN "drone_id" bigint NOT NULL, ADD COLUMN "medication_id" bigint NOT NULL, ADD
 CONSTRAINT "fk_orders_drone" FOREIGN KEY ("drone_id") REFERENCES "drones" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_orders_medication" FOREIGN KEY ("medication_id") REFERENCES "medications" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
