# GoRose-ORM-Pro for Commercial(专业版)

[![GoDoc](https://godoc.org/github.com/tobycroft/gorose-pro?status.svg)](https://godoc.org/github.com/tobycroft/gorose-pro)
[![Go Report Card](https://goreportcard.com/badge/github.com/tobycroft/gorose-pro)](https://goreportcard.com/report/github.com/tobycroft/gorose-pro)
[![GitHub release](https://img.shields.io/github/release/tobycroft/gorose.svg)](https://github.com/tobycroft/gorose-pro/releases/latest)
[![Gitter](https://badges.gitter.im/tobycroft/gorose.svg)](https://gitter.im/gorose/wechat)
![GitHub](https://img.shields.io/github/license/tobycroft/gorose?color=blue)
![GitHub All Releases](https://img.shields.io/github/downloads/tobycroft/gorose/total?color=blue)
<a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=P0R-T6lnM--WHzgvGPnbd58US3IUoDlW&jump_from=webapi">
94537310
<img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="gorose-orm" title="gorose-orm"></a>

~~~
 ██████╗  ██████╗ ██████╗  ██████╗ ███████╗███████╗    ██████╗ ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗██╔══██╗██╔═══██╗██╔════╝██╔════╝    ██╔══██╗██╔══██╗██╔═══██╗
██║  ███╗██║   ██║██████╔╝██║   ██║███████╗█████╗█████╗██████╔╝██████╔╝██║   ██║
██║   ██║██║   ██║██╔══██╗██║   ██║╚════██║██╔══╝╚════╝██╔═══╝ ██╔══██╗██║   ██║
╚██████╔╝╚██████╔╝██║  ██║╚██████╔╝███████║███████╗    ██║     ██║  ██║╚██████╔╝
 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝    ╚═╝     ╚═╝  ╚═╝ ╚═════╝ 
~~~

## 原版和Pro版本区别（原版没有的功能）+（猜你关心）

- 直觉式编程，更加适合ThinkPHP/Laravel开发人员
- 支持事务嵌套
- go get -u 直接升级，每次升级均做到向上向下兼容无需担心更新后不兼容导致的事故
- 跟深度支持MySQL和MariaDB
- 文档超详细，你可以在我的文档中找到99%
- 采用"直觉编程"优化，即使没用过也能更快上手
- 100%兼容原版
- 更快的PR/BUG响应+修复速度
- 所有的更新/Bug修复完全来自于当前正在编写的商业项目，不可能出现更新后不管的情况
- 支持复杂Where/OrWhere条件下的and/or条件查询（复杂环境下极好用！）

## 本项目目的

- 为了能让开发者更加习惯直觉式开发，
- 为了解决原版框架不再更新维护后可能带来的风险问题
- 为了解决原框架在商业项目实战中出现的各类弊端
- 巨细无遗的文档，无论你是PHP->Go还是纯新手，你都可以在文档中找到对应

## 故障修复

- 修复了高并发下，where等参数的的脏数据问题(如果你在用原版，避免生产环境使用单db)
- 修复了Paginate不能用的问题，并且新增Paginator，让返回更加清晰

## 实例文档（Wiki）

- 增删改
    - [增加Insert](https://github.com/tobycroft/gorose-pro/wiki/Insert新增数据)
    - [删除Delete](https://github.com/tobycroft/gorose-pro/wiki/Delete删除数据)
    - [更新Update](https://github.com/tobycroft/gorose-pro/wiki/Update方法)
- 单条查询（对象）(Map[string]interface{})
    - [Find/First返回对象](https://github.com/tobycroft/gorose-pro/wiki/Find-First查询返回Obj对象方法)
- 多条/联合查询（[]Map[string]interface{}
    - [Get/Select返回数组](https://github.com/tobycroft/gorose-pro/wiki/Get-Select方法)
    - [Join联合查询](https://github.com/tobycroft/gorose-pro/wiki/Join-Select方法)
    - [Paginator复杂的子查询分页构建](https://github.com/tobycroft/gorose-pro/wiki/Paginator复杂的子查询分页构建)
- Query方法
    - [Query方法使用原生语句查询](https://github.com/tobycroft/gorose-pro/wiki/Query方法)
- 嵌套事务
    - [支付环境下复杂的嵌套事务实例](https://github.com/tobycroft/gorose-pro/wiki/支付环境下复杂的嵌套事务)

## 简介

GorosePro是一个GolangOrm升级改版项目，在支持原框架所有功能的基础上修复了BUG，更加适合复杂的商业项目

支持解耦式开发和直觉式编程，大大降低你的试错成本，让小型项目开发速更快，让大型项目更加容易维护

## 安装

- go.mod 中添加

```bash
require github.com/tobycroft/gorose-pro v1.2.5
```

- go get

```bash
go get -u github.com/tobycroft/gorose-pro
```

## 支持驱动

- mysql : https://github.com/go-sql-driver/mysql
- sqlite3 : https://github.com/mattn/go-sqlite3
- postgres : https://github.com/lib/pq
- oracle : https://github.com/mattn/go-oci8
- mssql : https://github.com/denisenkom/go-mssqldb
- clickhouse : https://github.com/kshvakov/clickhouse

```go
db.Table("table_name").Fields().Where().GroupBy().Having().OrderBy().Limit().Select()
db.Table(&ModelStruct).Data().Replace()
db.Table(&ModelStruct).Data().Insert()
db.Table(....).Data().Where().Update()
db.Table(....).Where().Delete()
```

## 配置和链接初始化

简单配置DSN

```go
var conf gorose.Config
conf.Driver = "mysql"
conf.SetMaxIdleConns = 90
conf.SetMaxOpenConns = 300
conf.Prefix = ""
conf.Dsn = dsn()
return &conf
```

更多配置, 可以配置集群,甚至可以同时配置不同数据库在一个集群中, 数据库会随机选择集群的数据库来完成对应的读写操作, 其中master是写库, slave是读库, 需要自己做好主从复制, 这里只负责读写

```go
var config1 = gorose.Config{Dsn: 上面的dsn}
var config2 = gorose.Config{Dsn:  上面的dsn}
var config3 = gorose.Config{Dsn:  上面的dsn}
var config4 = gorose.Config{Dsn:  上面的dsn}
var configCluster = &gorose.ConfigCluster{
Master:  []gorose.Config{config3, config4},
Slave: []gorose.Config{config1, config2},
Driver: "sqlite3",
}
```

初始化使用

```go
var engin *gorose.Engin
engin, err := Open(config)
//engin, err := Open(configCluster)

if err != nil {
panic(err.Error())
}
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/tobycroft/gorose-pro.svg)](https://starchart.cc/tobycroft/gorose-pro)
