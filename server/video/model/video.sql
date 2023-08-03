CREATE TABLE `video` (
                         `id` bigint(20) NOT NULL,
                         `user_id` bigint(20) NOT NULL,
                         `title` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `play_url` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `cover_url` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `delete_time` datetime DEFAULT NULL,
                         `video_sha256` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `id` (`id`),
                         UNIQUE KEY `video_sha256` (`video_sha256`),
                         KEY `user_id__index` (`user_id`),
                         KEY `video_sha256__index` (`video_sha256`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci

