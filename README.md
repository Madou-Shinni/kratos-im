## kratos-im

## Teacher

- [x] [Go微服务系统精讲 Go-Zero全流程实战即时通讯 ( IM ) by 木兮](https://coding.imooc.com/class/826.html)

## Description

- [x] [Go-Kratos官方文档](https://go-kratos.dev/en/docs/getting-started/start/)

使用kratos构建的im系统，支持单聊、群聊、消息已读、好友列表、群成员列表、在线状态等功能。

## 架构图

![kratos-im.png](assets%2Fkratos-im.png)

## Quick Start

依赖组件

1. [x] go1.21
2. [x] mysql8.0
3. [x] redis6.0
4. [x] kafka2.8.1
5. [x] etcd3.5.0
6. [x] mongodb5.0

### 1. 安装依赖

```shell
go mod tidy

make init
```

### 2. 启动服务

```shell
kratos run
```

## 客户端

努力开发中...