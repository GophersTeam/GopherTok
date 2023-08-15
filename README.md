# GopherTok

## 架构

...

## 🚀技术栈

| 功能               | 实现                                   |
| :----------------- | -------------------------------------- |
| http框架           | go-zero                                |
| rpc框架            | go-zero                                |
| orm框架            | gorm                                   |
| 数据库             | Innodb-cluster,redis-cluster,mongodb   |
| 对象存储           | 腾讯云cos,minio                        |
| 服务发现与配置中心 | etcd,nacos                             |
| 链路追踪           | jaeger                                 |
| 服务监控           | prometheus,grafana                     |
| 消息队列           | kafka                                  |
| 日志搜集           | filebeat,go-stash,elasticsearch,kibana |
| 网关               | traefik                                |
| 部署               | Docker,docer-compose                   |
| CI/CD              | Github Action                          |

## 高可用

* mysql选择`innodb-cluster`

![image-20230814172330152](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230814172330152.png)



* redis选择`redis-cluster`

![在这里插入图片描述](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3lyeDQyMDkwOQ==,size_16,color_FFFFFF,t_70.png)

* minio集群

![MinIO分布式集群架构](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/a36949e0b971475499fd9ec95ad3b32d~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0-20230718162200891-20230814172546027.awebp)

