CREATE TABLE `users` (
  `id` bigserial PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `hashed_password` varchar(255) NOT NULL
);

CREATE TABLE `books` (
  `id` bigserial PRIMARY KEY,
  `title` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `ISBN` varchar(255) UNIQUE NOT NULL,
  `quantity` bigint NOT NULL,
  `category` varchar(255) NOT NULL
);

CREATE TABLE `lending_records` (
  `id` bigserial PRIMARY KEY,
  `book` bigint NOT NULL,
  `borrower` bigint NOT NULL,
  `is_return` boolean NOT NULL DEFAULT false,
  `borrow_date` timestamptz NOT NULL DEFAULT (now()),
  `return_date` timestamptz NOT NULL DEFAULT (0001-01-01 00:00:00 +0000 UTC)
);

ALTER TABLE `lending_records` ADD FOREIGN KEY (`book`) REFERENCES `books` (`id`);

ALTER TABLE `lending_records` ADD FOREIGN KEY (`borrower`) REFERENCES `users` (`id`);
