# DnsSearch

---
#### 功能描述

>根据IP获取DNS解析记录
---
#### 注意事项

>默认10个go程，默认不设置代理，本机跑太快会很快封ip，设置go程数小于10！！！
---
#### 代理设置
##### redis读取
>代理池：https://github.com/lennon-liu/proxy_pool.git
``` 
-pf 127.0.0.1：6379 -pk proxypool
```
---
##### 文件读取：

```
-pf proxypools.txt
```
##### 文件内容format：

```
180.101.159.?:5000
180.101.159.?:5000
```
---
#### 编译

```
git clone https://github.com/lennon-liu/ip-query-domain.git
cd dns_search/cd 
go build -i dbs_search cmd.go
```
---
#### 参数

```
  -i string
        input file path or os.stdin
  -n int
        proxy Tolerant (default 5)//ip对应域名历史记录数
  -o string
        output file path or os.stdout
  -pa string
        proxytool IpAdddress fmt: 127.0.0.1:6379
  -pf string
        proxytool file path
  -pk string
        proxytool key fmt: proxytool
  -s int
        request Scanners (default 10) 扫描go程数
  -t int
        request timeout (default 2) 请求延迟
```
#### 快速使用

```
./dbs_search -i input.txt
./dbs_search -i input.txt -o output.txt
./dbs_search -i input.txt -o output.txt -n 1 -s 10 -t 2
./dbs_search -i input.txt -pf 127.0.0.1：6379 -pk proxypool
./dbs_search -i input.txt -pf proxypools.txt
```

#### 博客

>https://lennon.work/