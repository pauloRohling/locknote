CREATE TABLE "notes"
(
    id         uuid PRIMARY KEY      DEFAULT uuid_generate_v7(),
    title      varchar(255) NOT NULL,
    content    varchar      NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT now(),
    created_by uuid         NOT NULL,

    CONSTRAINT "notes_uuid_v7" CHECK (uuid_extract_version(id) = 7),
    CONSTRAINT "notes_created_by_v7" CHECK (uuid_extract_version(created_by) = 7)
);

CREATE INDEX "notes_created_by_idx" ON "notes" ("created_by");