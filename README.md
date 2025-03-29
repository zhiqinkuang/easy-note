# Easy Note

*  [原项目参考链接](https://github.com/cloudwego/biz-demo/tree/main/easy_note)
## 项目描述
* 一个用来学习微服务的小案例,实现 note 和用户功能
* 使用的框架是 kitex + Hertz 
* 系统框架是 etcd 用于服务发现
* 具体的框架结构如图所示

![系统框架](https://pic-bed-private.oss-cn-shanghai.aliyuncs.com/pic-bedeasy-note-arch.png)

## 项目实现思路

### 项目初始化
* 创建docker 需要的环境
* 建立idl文件夹,规定接口建立 note.thrift, user.thrift, api.thrift
* 建立Makefile, 写入代码生成指令
* 建立cmd文件夹,分别创建note,user,api子文件夹用来存放kitex的代码,建立Makefile,设置指令,用里面的代码生成初始化指令
* 在cmd/note or cmd/user 文件下面的创建  pake、dal、rpc、serivce 文件夹用来放文件
* 在 cmd/api 里面创建 mw(中间件)、rpc
* 在项目根目录下安装 hertz 之后执行 go mod tidy

### 项目依赖
* 安装项目的依赖包
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm/logger
go get -u gorm.io/plugin/opentelemetry/logging/logrus
go get -u gorm.io/plugin/opentelemetry/tracing
go get github.com/cloudwego/kitex/pkg/endpoint
go get github.com/cloudwego/kitex/pkg/klog
go get github.com/cloudwego/kitex/pkg/rpcinfo
```



