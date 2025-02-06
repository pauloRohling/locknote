CREATE TABLE "users"
(
    id         uuid PRIMARY KEY     DEFAULT uuid_generate_v7(),
    name       varchar(50) NOT NULL,
    email      varchar     NOT NULL,
    password   varchar     NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    created_by uuid        NOT NULL,

    CONSTRAINT "uuid_users_uuid_v7" CHECK (uuid_extract_version(id) = 7),
    CONSTRAINT "created_by_users_uuid_v7" CHECK (uuid_extract_version(created_by) = 7)
);

CREATE UNIQUE INDEX "users_email_idx" ON "users" ("email");