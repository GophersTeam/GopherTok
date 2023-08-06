CREATE TABLE `follow_subject`IF NOT EXISTS (
                                  `id` int NOT NULL COMMENT '唯一id',
                                  `user_id` int DEFAULT NULL COMMENT '被关注用户的id',
                                  `follower_id` int DEFAULT NULL COMMENT '关注者id',
                                  `is_follow` varchar(255) DEFAULT NULL COMMENT '是否关注',
                                  `creat_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;