CREATE DATABASE rin;
GRANT ALL ON rin.* TO 'rin'@'%';

USE rin;
CREATE TABLE account
(
    id           varchar(36) PRIMARY KEY,
    name         varchar(16) UNIQUE,
    display_name varchar(32),
    email        varchar(64) UNIQUE,
    password     varchar(128),
    created_at   datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at   datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);