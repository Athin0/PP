CREATE SCHEMA IF NOT EXISTS seq;

CREATE TABLE IF NOT EXISTS seq.storage
(
    title varchar(256) PRIMARY KEY,
    data  jsonb NOT NULL
);
