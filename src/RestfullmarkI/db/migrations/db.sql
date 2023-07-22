CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `code` integer,
  `created_at` timestamp,
  `update_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `activities` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `title` varchar(255),
  `created_at` timestamp,
  `update_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `todolist` (
  `id` integer PRIMARY KEY,
  `activities_id` integer,
  `title` varchar(255),
  `status` enum,
  `created_at` timestamp,
  `update_at` timestamp,
  `deleted_at` timestamp
);

ALTER TABLE `activities` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `todolist` ADD FOREIGN KEY (`activities_id`) REFERENCES `activities` (`id`);
