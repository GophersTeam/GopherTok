# GopherTok

English | [简体中文](README-cn.MD)

> **Project introduction**：http://gophertok.kuaji.club/
>
> **Project presentation**：https://www.bilibili.com/video/BV1sh4y1N7bq/?spm_id_from=333.999.0.0

| <img src="https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308222108323.png" alt="{09951663-C990-6AA2-14C8-28D9C1DDBDCD}" style="zoom: 25%;" /> | The sixth Bytedance Youth training camp big project works, a simple version of Tiktok project ，built with  go-zero  microservice . Completed by the [gopher team](https://github.com/GophersTeam/GopherTok) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

## 🪵Architecture

![eb4302aa8c255a470e8be4becfda63ad](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103203.png)

## 🚀Technology stack

| Feature                                                   | **Implementation**                     |
| :-------------------------------------------------------- | -------------------------------------- |
| HTTP framework                                            | go-zero                                |
| RPC framework                                             | go-zero                                |
| ORM framework                                             | gorm                                   |
| Database                                                  | Innodb-cluster,redis-cluster,mongodb   |
| Object storage                                            | Tencent Cloud COS, Minio               |
| Service discovery, registration, and configuration center | etcd,nacos                             |
| Tracing                                                   | jaeger                                 |
| Service monitoring                                        | prometheus,grafana                     |
| Message queue                                             | kafka                                  |
| Log collection                                            | filebeat,go-stash,elasticsearch,kibana |
| Gateway                                                   | traefik                                |
| Deployment                                                | Docker,docker-compose                  |
| CI/CD                                                     | Github Action                          |

## High availability

*   `innodb-cluster`

![image-20230814172330152](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103716.png)

*  `redis-cluster`

![在这里插入图片描述](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103244.png)

* `minio cluster`

![MinIO分布式集群架构](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103895.awebp)

* 4 nodes

![image-20230816101826428](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103388.png)

* `kafka cluster`

![image-20230816101130893](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103086.png)

* 3rd node

![image-20230816101331794](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103560.png)

* The APIs and RPCs of the services start multiple docker instances. Traefik is used to load balance the APIs, and the RPCs implement load balancing via etcd to ensure service reliability and easy scalability during peak periods.

## High concurrency

The RPC calls of each service are executed concurrently, significantly reducing response time. `Redis` is used as a cache for high-frequency data, reducing the pressure on `MySQL`. `Kafka` is used to asynchronously write to `MySQL`, increasing system throughput.·

## High performance

`kafka` uses clustered writes to greatly `reduce` disk io and network io

### Configuration center and service discovery/registry center

`Nacos` is used as the configuration center

![image-20230818163632603](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818163632603.png)

`Etcd` is used for service discovery and registry

![e45ceb303cceb5ea188b8fa11f66c768](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/e45ceb303cceb5ea188b8fa11f66c768.png)

### Tracing

`Jaeger` is used for link tracing across services.

![796364212238fb72b302c76a95f124b1](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103643.png)

###  Log collection

![image-20230818164131821](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818164131821.png)

`Filebeat` collects logs and sends them to `Kafka` for buffering. `Go-stash` filters the logs based on configuration and outputs the filtered logs to `Elasticsearch`. `Kibana` is responsible for visualizing the logs.

![39ca160fbd2b2b385622deef2e79ba03](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103804.png)

### Supervisory control

`Prometheus` is used for service monitoring, with `Grafana` providing a visualization interface.

![42ba4597865261dcddcd1545d78c3d4f](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103873.png)

visualized by `grafana`

![image-20230818160820149](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230818160820149.png)

### Gateway

`Traefik` is used as the gateway, load balancing requests to the API containers of each service based on routing rules.

![image-20230818163032128](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103246.png)

load-balanced to various service 'api' container instances

![image-20230818164454219](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103425.png)

## CI/CD

* Use `GitHub Action` for `CI/CD` and automate testing after each commit
* Then you can manually build the image of each service, and automatically push it to the `dockerhub` after it is built
* Then ssh into the remote server, using the new image and already written 'docker-compose' automatically deploy the new container

## 🎉Thanks

|   [Bytedance youth Camp](https://youthcamp.bytedance.com/)   |
| :----------------------------------------------------------: |
| <img src="https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202308250103647.webp" alt="青训营" style="zoom: 67%;" /> |

## 📣Licence

**GopherTok** is open-source under the **MIT** license. Please follow the [MIT open source certificate](https://github.com/MashiroC/begonia/blob/master/LICENSE) rules under the premise of use

