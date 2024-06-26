-- Modify "orders" table
ALTER TABLE "orders" DROP CONSTRAINT "chk_orders_quantity", DROP CONSTRAINT "fk_orders_medication", ALTER COLUMN "drone_id" DROP NOT NULL, DROP COLUMN "medication_id", ADD COLUMN "medication_code" text NULL, ADD
 CONSTRAINT "fk_orders_medication" FOREIGN KEY ("medication_code") REFERENCES "medications" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION;
