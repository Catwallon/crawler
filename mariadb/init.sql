CREATE DATABASE IF NOT EXISTS db;
USE db;

CREATE TABLE `db`.`pages` (
	`website` TEXT NOT NULL,
	`url` TEXT NOT NULL,
	`title` TEXT NOT NULL,
	`description` TEXT NOT NULL,
	`keywords` JSON NOT NULL,
	`lang` TEXT NOT NULL
) ENGINE = InnoDB; 