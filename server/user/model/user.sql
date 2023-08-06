CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL,
                        `username` varchar(32) COLLATE utf8mb4_unicode_520_ci NOT NULL,
                        `password` varchar(32) COLLATE utf8mb4_unicode_520_ci NOT NULL,
                        `avatar` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
                        `background_image` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
                        `signature` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
                        `follow_count` bigint(20) NOT NULL DEFAULT '0',
                        `follower_count` bigint(20) NOT NULL DEFAULT '0',
                        `friend_count` bigint(20) NOT NULL DEFAULT '0',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `id` (`id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci

