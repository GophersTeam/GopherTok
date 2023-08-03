CREATE TABLE `video` (
                         `id` bigint(20) NOT NULL,
                         `user_id` bigint(20) NOT NULL,
                         `title` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
                         `play_url` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `cover_url` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `create_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
                         `update_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
                         `delete_time` datetime DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `id` (`id`),
                         KEY `user_id__index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci

