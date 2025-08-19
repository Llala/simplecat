CREATE TABLE "applications" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "source_text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "source_unit" (
  "id" int PRIMARY KEY,
  "application_id" int NOT NULL,
  "translation_unit_id" int NOT NULL,
  "text" varchar
);

CREATE TABLE "translation_unit" (
  "id" int PRIMARY KEY,
  "application_id" int NOT NULL,
  "source_unit_id" int NOT NULL,
  "text" varchar
);

ALTER TABLE "source_unit" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");

ALTER TABLE "source_unit" ADD FOREIGN KEY ("translation_unit_id") REFERENCES "translation_unit" ("id");

ALTER TABLE "translation_unit" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");

ALTER TABLE "translation_unit" ADD FOREIGN KEY ("source_unit_id") REFERENCES "source_unit" ("id");