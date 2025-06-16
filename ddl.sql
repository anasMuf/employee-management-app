CREATE DATABASE hacktiv8_p2_gc1;

USE hacktiv8_p2_gc1;

CREATE TABLE employees (
	id INT auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	phone varchar(100) NOT NULL,
	created_at DATETIME NULL,
	updated_at DATETIME NULL,
	CONSTRAINT employees_pk PRIMARY KEY (id),
	CONSTRAINT employees_unique_email UNIQUE KEY (email),
	CONSTRAINT employees_unique_phone UNIQUE KEY (phone)
);