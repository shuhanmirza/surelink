CREATE TABLE "redirection_map"
(
    "uuid"       varchar PRIMARY KEY NOT NULL,
    "url"        varchar             NOT NULL,
    "created_at" timestamptz         NOT NULL DEFAULT (now())
);
