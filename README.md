# Fabric-sdk-go方式获取区块链概况

## 简介

通过fabric-sdk-go方式获取区块链网络的概要信息，高度，哈希等，由于go语言方式的sdk提供的接口和java/node方式的接口有较大的差异，网上资料也比较匮乏，所以整理上传。

## 使用步骤

1. 编辑sdk-config.yaml配置文件
2. go run main.go运行，即可输入如下日志

## 易错点提醒

- 本机测试完全没有问题
- 远程主机测试，比较有问题
- SDK使用的配置文件与连接文件有差异，不能照搬

## 日志

```
❯ go run main.go
2021/03/17 14:24:59 ---> 创建SDK成功
区块高度:
7
当前区块Hash:
0e243e24b8f7a5c394b8ae376bdf3abf31be101056e533a0ab044eb7149500c7
前一区块Hash:
ec8ad99789a48e076671c976ef5ff5f9e69023f6dcf25137d5bb15c1ce17c723
区块编号: 6
区块Hash:
fa68edb07022b971fc34c31b22f582366567360ff3e662a692a11c62318a0c3a
区块编号: 6
区块Hash:
fa68edb07022b971fc34c31b22f582366567360ff3e662a692a11c62318a0c3a
通道名称: mychannel
区块个数: 2
锚节点:
  主机:peer0.org1.example.com
  端口:7051
  机构:Org1MSP
```
