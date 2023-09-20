CREATE TABLE `favor` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `Userid` bigint(20) DEFAULT NULL,
                         `Videoid` bigint(20) DEFAULT NULL,
                         `Status` int(11) DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `user_id__index` (`Userid`),
                         KEY `video_id__index` (`Videoid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

