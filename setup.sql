CREATE DATABASE `hamlistings` /*!40100 COLLATE 'utf8_general_ci' */;


USE `hamlistings`;

CREATE TABLE `user` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(30) NOT NULL,
	`email` VARCHAR(256) NOT NULL,
	`password` VARCHAR(128) NOT NULL,
	`reg_key` VARCHAR(60) NOT NULL,
	`reset_key` VARCHAR(60) NOT NULL,
	`email_verified` TINYINT(1) default 0,
	`created` DATETIME NOT NULL,
	`updated` DATETIME NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `username` (`username`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1000;

CREATE TABLE `listing` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`user_id` INT(11) NOT NULL,
	`txt` VARCHAR(4096) NOT NULL,
	`condition` VARCHAR(100) NOT NULL,
	`item_name` VARCHAR(100) NOT NULL,
	`payment_type` VARCHAR(100) NOT NULL,
	`price` float(11) NOT NULL,
	`contact_method` VARCHAR(100) NOT NULL,
	`contact` VARCHAR(100) NOT NULL,
	`delivery_method` VARCHAR(100) NOT NULL,
	`category` TINYINT(2) NOT NULL,
	`archived` TINYINT(1) NOT NULL,
	`created` DATETIME NOT NULL,
	`updated` DATETIME NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `FK__user` (`user_id`),
	CONSTRAINT `FKuser` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1000;
