-- Modify "drone_medications" table
ALTER TABLE "drone_medications" DROP CONSTRAINT "drone_medications_pkey", ADD PRIMARY KEY ("medication_id", "drone_id");
