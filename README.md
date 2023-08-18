# GopherTok

## 架构

...

## 🚀技术栈

| 功能                     | 实现                                   |
| :----------------------- | -------------------------------------- |
| http框架                 | go-zero                                |
| rpc框架                  | go-zero                                |
| orm框架                  | gorm                                   |
| 数据库                   | Innodb-cluster,redis-cluster,mongodb   |
| 对象存储                 | 腾讯云cos,minio                        |
| 服务发现、注册与配置中心 | etcd,nacos                             |
| 链路追踪                 | jaeger                                 |
| 服务监控                 | prometheus,grafana                     |
| 消息队列                 | kafka                                  |
| 日志搜集                 | filebeat,go-stash,elasticsearch,kibana |
| 网关                     | traefik                                |
| 部署                     | Docker,docer-compose                   |
| CI/CD                    | Github Action                          |

## 高可用

* mysql选择`innodb-cluster`

![image-20230814172330152](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230814172330152.png)



* redis选择`redis-cluster`

![在这里插入图片描述](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3lyeDQyMDkwOQ==,size_16,color_FFFFFF,t_70.png)

* minio集群

![MinIO分布式集群架构](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/a36949e0b971475499fd9ec95ad3b32d~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0-20230718162200891-20230814172546027.awebp)

四个节点

![image-20230816101826428](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230816101826428.png)

* kafka集群

![image-20230816101130893](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230816101130893.png)

3节点

![image-20230816101331794](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230816101331794.png)

* 服务的api和rpc启动多个docker实例，api使用traefik负载均衡，rpc通过etcd实现负载，保证服务的可靠性，高峰时期可以轻松扩容

## 高并发

调用各个服务的rpc时采用并发调用，大大增加系统的吞吐量

## 高性能

多处采用redis作缓冲，减少了mysql压力，各个服务使用kafka异步写入mysql，减少了响应时间

## 链路追踪



![796364212238fb72b302c76a95f124b1](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/796364212238fb72b302c76a95f124b1.png)

## 日志搜集



![39ca160fbd2b2b385622deef2e79ba03](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/39ca160fbd2b2b385622deef2e79ba03.png)

## 监控



![42ba4597865261dcddcd1545d78c3d4f](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/42ba4597865261dcddcd1545d78c3d4f.png)

![image-20230818160820149](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818160820149.png)

## 网关

![image-20230818163032128](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818163032128.png)

## CI/CD

使用Github Action进行CI/CD，每次提交上去后进行自动化测试，然后可以手动构建各个服务的镜像，构建好后自动推送到dockerhub上面，之后再ssh登录远程服务器，利用新的镜像和已经写好的docker-compose自动部署好新的容器
