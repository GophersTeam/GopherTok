# chat
## 技术选型
MySQL+Redis

## 数据库设计
消息表

| 字段         | 数据类型 | 说明     |
| ------------ | -------- | -------- |
| id           | bigint   | 主键id   |
| from_user_id | bigint   | 发送方id |
| to_user_id   | bigint   | 接收方id |
| content      | varchar  | 消息内容 |
| create_time  | bigint   | 发送时间 |


## 亮点
- 使用redis的map结构缓存用户的好友列表之间的最后一条消息，提高系统的吞吐量
- 使用雪花算法生成消息id，在分布式环境下保证消息id的唯一性