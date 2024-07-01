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
) ;

 CREATE TABLE addresses (
    id INT auto_increment PRIMARY KEY,
    street VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    zipcode VARCHAR(20) NOT NULL,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;

CREATE TABLE user_addresses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    address_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (address_id) REFERENCES addresses(id),
    UNIQUE KEY user_address_unique (user_id, address_id)
)ENGINE=INNODB;