CREATE DATABASE account;
GRANT ALL ON account.* TO 'rin'@'%';

USE account;
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