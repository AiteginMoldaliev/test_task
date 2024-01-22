CREATE TABLE "persons" (
	"id" bigserial PRIMARY KEY,
    "firstname" varchar NOT NULL,
	"surname" varchar NOT NULL,
	"patronymic" varchar NULL,
	"gender" varchar NOT NULL,
	"age" bigint NOT NULL,
	"nationality" varchar NOT NULL,
    "created_at" TIMESTAMPTZ not null DEFAULT (now())
);

CREATE INDEX ON "persons" ("firstname");

CREATE INDEX ON "persons" ("surname");