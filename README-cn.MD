# GopherTok

> 项目介绍：http://gophertok.kuaji.club/
>

| <img src="https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308222108323.png" alt="{09951663-C990-6AA2-14C8-28D9C1DDBDCD}" style="zoom: 25%;" /> | 第六届字节跳动青训营大项目作品，使用go-zero搭建的微服务项目，由[gopher小队](https://github.com/GophersTeam/GopherTok)完成，一个简易版抖音项目 |
| ------------------------------------------------------------ | ------------------------------------------------------------ |



## 🪵架构

![eb4302aa8c255a470e8be4becfda63ad](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308230226461.png)

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

* 服务的`api`和`rpc`启动多个do`cker实例，api使用`traefik`负载均衡，`rpc`通过`etcd`实现负载，保证服务的可靠性，高峰时期可以轻松扩容

## 高并发

* 调用各个服务的`rpc`时采用并发调用，显著减少了响应时间
* 将高频率的数据采用`redis`作缓存，减少了`mysql`压力

* 使用`kafka`异步写入`mysql`，增加系统吞吐量

## 高性能

* kafka采用聚集写入，大大减少磁盘io和网络io

## 配置中心和服务发现、注册中心

nacos作为配置中心

![image-20230818163632603](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818163632603.png)

etcd作为服务发现和注册中心

![e45ceb303cceb5ea188b8fa11f66c768](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/e45ceb303cceb5ea188b8fa11f66c768.png)

## 链路追踪

使用`jaeger`作为链路追踪

![796364212238fb72b302c76a95f124b1](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/796364212238fb72b302c76a95f124b1.png)

## 日志搜集

![image-20230818164131821](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818164131821.png)

`filebeat`收集业务日志，然后将日志输出到`kafka`中作为缓冲，`go-stash`获取`kafka`中日志根据配置过滤字段，然后将过滤后的字段输出到`elasticsearch`中，最后由`kibana`负责呈现日志

![39ca160fbd2b2b385622deef2e79ba03](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/39ca160fbd2b2b385622deef2e79ba03.png)

## 监控

使用`prometheus`进行服务监控

![42ba4597865261dcddcd1545d78c3d4f](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/42ba4597865261dcddcd1545d78c3d4f.png)

再由`grafana`进行可视化呈现

![image-20230818160820149](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818160820149.png)



## 网关

使用`traefik`作为网关，根据路由规则将请求负载均衡到各个服务的`api`上面

![image-20230818163032128](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818163032128.png)

再通过负载均衡到各个服务`api`容器实例

![image-20230818164454219](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818164454219.png)

## CI/CD

* 使用`Github Action`进行CI/CD，每次提交上去后进行自动化测试
* 然后可以手动构建各个服务的镜像，构建好后自动推送到`dockerhub`上面
* 再ssh登录远程服务器，利用新的镜像和已经写好的`docker-compose`自动部署好新的容器

## 🎉感谢

|      [字节跳动青训营](https://youthcamp.bytedance.com/)      |
| :----------------------------------------------------------: |
| <img src="https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308230232085.webp" alt="青训营" style="zoom: 67%;" /> |



## 📣许可证

**GopherTok** 在 **MIT** 许可证下开源,请在遵循 [MIT 开源证书](https://github.com/MashiroC/begonia/blob/master/LICENSE) 规则的前提下使用
