CREATE DATABASE IF NOT EXISTS db;
USE db;

CREATE TABLE `db`.`pages` (
	`website` TEXT NOT NULL,
	`url` TEXT NOT NULL,
	`title` TEXT NOT NULL,
	`description` TEXT NOT NULL,
	`keywords` TEXT NOT NULL,
	`score` INT NOT NULL
) ENGINE = InnoDB; 