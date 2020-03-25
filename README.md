# grpc-stream
An example of gRPC stream

## usage
- start server

```
docker run --rm -it -p 8080:8080 registry.cn-hangzhou.aliyuncs.com/knative-sample/server:2020-03-24_173752
```

- start client

```
docker run --rm -it -e GRPC_CONCURRENT="10000" -e GRPC_SERVER_ADDR=${server_addr}:8080 registry.cn-hangzhou.aliyuncs.com/knative-sample/client:2020-03-24_212051
```
  - GRPC_CONCURRENT 指定并发数
  - GRPC_SERVER_ADDR 指定 server 地址

```
└─# docker run --rm -it -e GRPC_SERVER_ADDR=30.5.123.238:8080 client
2020/03/24 13:59:52 resp: key: Client Stream List for [8]  --> resp index:8, value: 1
2020/03/24 13:59:52 resp: key: Client Stream List for [16]  --> resp index:16, value: 1
2020/03/24 13:59:52 resp: key: Client Stream List for [12]  --> resp index:12, value: 1
2020/03/24 13:59:52 resp: key: Client Stream List for [14]  --> resp index:14, value: 1
2020/03/24 13:59:52 resp: key: Client Stream List for [93]  --> resp index:93, value: 1
2020/03/24 13:59:52 resp: key: Client Stream List for [92]  --> resp index:92, value: 1
... ... 
```
