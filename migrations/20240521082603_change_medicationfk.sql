-- Modify "orders" table
ALTER TABLE "orders" DROP CONSTRAINT "fk_orders_medication", DROP COLUMN "medication_code", ADD COLUMN "medication_id" bigint NULL, ADD
 CONSTRAINT "fk_orders_medication" FOREIGN KEY ("medication_id") REFERENCES "medications" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
