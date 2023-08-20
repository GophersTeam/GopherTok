# video

## 发布视频

> 支持三种方式存储，minio集群，腾讯云cos，本地，可通过nacos配置

### 流程

1. 首先计算出视频的sha256值，在redis中查询是否存在该值，若存在说明之前用户传过一样的视频，则直接从redis查询视频的url和封面的url,返回上传成功
2. 将计算出的sha256存入redis中
3. 拼接好文件路径
4. 使用ffmpeg截取视频第一帧作为封面
5. 根据三种存储方式对应上传，并生成相应的url
6. 利用雪花算法生成id，带上视频的基本信息写入kafka中
7. 返回视频上传成功
8. kafka消费数据，写入mysql和redis

## 查看用户发布的全部视频

1. 根据user_id从mysql中查询全部视频
2. 并发调用其他rpc拼接最终返回信息

## 视频流

1. 根据latest_time从mysql查询30个视频
2. 并发调用其他rpc拼接最终返回信息

## mq

* kafka聚集写入

![image-20230708153904061](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230708153904061.png)

之前每向kafka发送一条消息就会产生一次网络 IO 和一次磁盘 IO，做消息聚合后，比如聚合 100 条消息后再发送给 Kafka，这个时候 100 条消息才会产生一次网络 IO 和磁盘 IO，这样大大提高 Kafka 的吞吐和性能。并且有聚合时间兜底，就算消息数量达不到聚合要求，超过聚合最大时间也会聚合当前所有消息发送给Kafka

* 并发消费

![img](https://cdn.learnku.com/uploads/images/202207/15/73865/pBHJlrB3rC.webp!large)

通过多个 goroutine 来并行的消费数据

## 亮点

* 计算视频的sha256值可以做到秒传视频
* 利用ffmpeg截取视频第一帧作为封面
* 可以选择三种存储方式｜腾讯云cos|minio集群｜本地
* kafka异步写入视频信息已经相关优化
* 并发调用其他服务rpc拼接返回信息
* 并发写入数据库

## 参考

https://learnku.com/articles/69754