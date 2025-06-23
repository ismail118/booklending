CREATE TABLE `users` (
                         `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                         `name` varchar(255) NOT NULL,
                         `email` varchar(255) UNIQUE NOT NULL,
                         `hashed_password` varchar(255) NOT NULL
);

CREATE TABLE `books` (
                         `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                         `title` varchar(255) NOT NULL,
                         `author` varchar(255) NOT NULL,
                         `ISBN` varchar(255) UNIQUE NOT NULL,
                         `quantity` bigint NOT NULL,
                         `category` varchar(255) NOT NULL
);

CREATE TABLE `lending_records` (
                                   `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                   `book` bigint NOT NULL,
                                   `borrower` bigint NOT NULL,
                                   `is_return` boolean NOT NULL DEFAULT false,
                                   `borrow_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   `return_date` TIMESTAMP NULL DEFAULT NULL
);

ALTER TABLE `lending_records` ADD FOREIGN KEY (`book`) REFERENCES `books` (`id`);

ALTER TABLE `lending_records` ADD FOREIGN KEY (`borrower`) REFERENCES `users` (`id`);
