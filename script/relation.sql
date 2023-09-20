CREATE TABLE `follow_subject` (
                                  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一id',
                                  `user_id` bigint(20) DEFAULT NULL,
                                  `follower_id` bigint(20) DEFAULT NULL,
                                  `is_follow` tinyint(1) DEFAULT NULL,
                                  `creat_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  PRIMARY KEY (`id`) USING BTREE,
                                  KEY `uer_id__index` (`user_id`),
                                  KEY `follower_id__index` (`follower_id`)
) ENGINE=InnoDB AUTO_INCREMENT=141 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC

