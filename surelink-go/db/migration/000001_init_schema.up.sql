CREATE TABLE "url_map"
(
    "uid"             varchar PRIMARY KEY NOT NULL,
    "url"             varchar             NOT NULL,
    "time_redirected" bigint                       DEFAULT 0,
    "created_at"      timestamptz         NOT NULL DEFAULT (now())
);
