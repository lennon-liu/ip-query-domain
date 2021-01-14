# ip-query-domain

---
简单的爬虫 https://site.ip138.com/ 获取 ip 对应 dns解析记录

---
速度

默认100go程，本机跑会很快封ip，需配置代理 否则设置go程数小于10

---

代理
---
#### redis读取

代理池：https://github.com/jhao104/proxy_pool.git

---
文件读取：
format：
```
`180.101.159.?:5000\n`
`180.101.159.?:5000\n`
`180.101.159.?:5000\n`
`180.101.159.?:5000\n`
`180.101.159.?:5000\n`
```
---
#### how to make

```
git clone https://github.com/lennon-liu/ip-query-domain.git
`cd ip-query-domain`
`cd cmd`
`go build -i cmd.go`
````

---
#### help
```
`Usage of /tmp/___go_build_cmd_go:`
  `-i signal`
    	`send signal to a master process: stop, quit, reopen, reload`
  `-n int`
    	`proxy Tolerant (default 5)`
  `-o signal`
    	`send signal to a master process: stop, quit, reopen, reload`
  `-pa string`
    	`proxytool IpAdddress fmt: 127.0.0.1:6379 (default "127.0.0.1:6379")`
  `-pf string`
    	`proxytool file path`
  `-pk string`
    	`proxytool key fmt: proxytool (default "proxytool")`
  `-t int`
`​    	request timeout (default 2)`
```
#### blog

​	https://lennon.work/