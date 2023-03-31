CREATE TABLE "url_map"
(
    "uid"       varchar PRIMARY KEY NOT NULL,
    "url"        varchar             NOT NULL,
    "created_at" timestamptz         NOT NULL DEFAULT (now())
);
