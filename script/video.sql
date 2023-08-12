CREATE TABLE `video` (
                         `id` bigint(20) NOT NULL,
                         `user_id` bigint(20) NOT NULL,
                         `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `delete_time` datetime DEFAULT NULL,
                         `video_sha256` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `id` (`id`),
                         UNIQUE KEY `video_sha256` (`video_sha256`),
                         KEY `idx_user_id_create_time` (`user_id`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci

