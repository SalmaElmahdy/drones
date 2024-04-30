-- Modify "medications" table
ALTER TABLE "medications" ADD COLUMN "name" text NOT NULL, ADD COLUMN "weight" numeric NOT NULL, ADD COLUMN "code" text NOT NULL, ADD COLUMN "image" text NOT NULL;
-- Create index "uni_medications_code" to table: "medications"
CREATE UNIQUE INDEX "uni_medications_code" ON "medications" ("code");
-- Create index "uni_medications_image" to table: "medications"
CREATE UNIQUE INDEX "uni_medications_image" ON "medications" ("image");
