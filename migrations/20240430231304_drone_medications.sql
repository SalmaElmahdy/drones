-- Create "drone_medications" table
CREATE TABLE "drone_medications" (
  "medication_id" bigint NOT NULL,
  "drone_id" bigint NOT NULL,
  PRIMARY KEY ("medication_id", "drone_id"),
  CONSTRAINT "fk_drone_medications_drone" FOREIGN KEY ("drone_id") REFERENCES "drones" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_drone_medications_medication" FOREIGN KEY ("medication_id") REFERENCES "medications" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
