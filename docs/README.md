YYTars介绍
=========================
> YYTars基于腾讯开源的tars项目，增加了GO语言和Protobuf的支持，并针对YY内部的现状做了本地化处理（监控，日志，权限管理等），同时增加了客户端SDK方便使用（android，iOS，pc）。

##### 1.YYTars的特性
* RPC服务框架，支持同步和异步调用；<br>
* 跨语言开发，支持GO，C++，Java；<br>
* 协议使用Protobuf编码；<br>
* 高性能，提供10w/s的处理能力；<br>
* 提供服务限流，容灾能力；<br>
* 运维管理一体化；<br>

##### 2.基础服务支持，方便新业务的快速接入，开发，部署了基础服务的支持：
* 基于HTTP的短连接通道，处理请求/响应消息；<br>
* 基于Web Socket的长连接通道，处理广播消息；<br>

##### 3.YYP兼容方案
* YYP提供的基础服务，通过增加代理来为YYTars提供服务；<br>
* HTTP和Thrift服务，可以直接访问；<br>

![](http://code.com/tars/docs/raw/90571a37c083e54dfe428a5759839c7e9f9080d1/images/yytars1.png)

使用说明
=========================
 
##### 1.名词定义
* Appname，应用名，微服务名，业务名称等。
* Service，进程名，一个appname里面可以有多个service。
* ServantObj，接口名，一个service里面可以有多个servantObj，每个servantObj有一个对外服务的端口。
* Function，RPC调用的方法名，一个servantObj里面可以有多个Function。

##### 2.后端服务开发流程
* 准备工作
 - Framework的安装
 - Service项目生成
* Protobuf协议定义
* 业务代码编写
* 单元测试
* 构建
* 测试
* 部署

##### 3.移动端接入流程

设计详解
=========================

##### 1.Framework架构
##### 2.监控报警
##### 3.远程日志
##### 4.配置管理
##### 5.协议格式
* 客户端SDK访问协议

![](http://code.com/tars/docs/raw/90571a37c083e54dfe428a5759839c7e9f9080d1/images/yytars2.jpg)

* 后端RPC调用协议

![](http://code.com/tars/docs/raw/90571a37c083e54dfe428a5759839c7e9f9080d1/images/yytars3.jpg)

* Tars头

```C++
//请求包体
struct RequestPacket<br>
{
    1  require short        iVersion;         //版本号
    2  optional byte        cPacketType;      //包类型
    3  optional int         iMessageType;     //消息类型
    4  require int          iRequestId;       //请求ID
    5  require string       sServantName;     //servant名字
    6  require string       sFuncName;        //函数名称
    7  require vector<byte> sBuffer;          //二进制buffer
    8  optional int         iTimeout;         //超时时间（毫秒）
    9  optional map<string, string> context;  //业务上下文
    10 optional map<string, string> status;   //框架协议上下文
};

//响应包体
struct ResponsePacket
{
    1 require short         iVersion;       //版本号
    2 optional byte         cPacketType;    //包类型
    3 require int           iRequestId;     //请求ID
    4 optional int          iMessageType;   //消息类型
    5 optional int          iRet;           //返回值
    6 require vector<byte>  sBuffer;        //二进制流
    7 optional map<string, string> status;  //协议上下文
    8 optional string       sResultDesc;    //结果描述
};
```
其他
=========================

##### 1.YYTars框架部署流程
##### 2.移动端项目架构设计参考
##### 3.Web端项目架构设计参考
##### 4.技术支持技术支持支持
* 李佳林（909010427，cell：18602023813）
* 黄灏（909015003，cell：13922115830）
* 卢保全（909012229，cell：13632196080）
* 陈智新（909014329，cell：13760727431）
* 胡涛（909013868，cell：13112268576）

