CREATE TABLE `favor_subject`IF NOT EXISTS (
                                               `id` int NOT NULL COMMENT '唯一id',
                                               `viode_id` int DEFAULT NULL COMMENT '被点赞的id',
                                               `user_id` int DEFAULT NULL COMMENT '点赞者id',
                                               `is_favor` varchar(255) DEFAULT NULL COMMENT '是否点赞',
                                               `creat_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;