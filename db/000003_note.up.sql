CREATE TABLE "notes"
(
    id         uuid PRIMARY KEY      DEFAULT uuid_generate_v7(),
    title      varchar(255) NOT NULL,
    content    varchar      NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT now(),
    created_by uuid         NOT NULL,

    CONSTRAINT "uuid_notes_uuid_v7" CHECK (uuid_extract_version(id) = 7),
    CONSTRAINT "created_by_notes_uuid_v7" CHECK (uuid_extract_version(created_by) = 7)
);