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