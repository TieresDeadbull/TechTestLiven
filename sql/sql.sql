CREATE DATABASE IF NOT EXISTS liven;
USE liven;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    passphrase varchar(100) not null,
    phoneNumber varchar(16) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;