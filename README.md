# GoRose-ORM-Pro 完全免费的MySQL数据库ORM

[![GoDoc](https://godoc.org/github.com/tobycroft/gorose-pro?status.svg)](https://godoc.org/github.com/tobycroft/gorose-pro)
[![Go Report Card](https://goreportcard.com/badge/github.com/tobycroft/gorose-pro)](https://goreportcard.com/report/github.com/tobycroft/gorose-pro)
[![GitHub release](https://img.shields.io/github/release/tobycroft/gorose-pro.svg)](https://github.com/tobycroft/gorose-pro/releases/latest)
![GitHub](https://img.shields.io/github/license/tobycroft/gorose-pro?color=blue)
![GitHub All Releases](https://img.shields.io/github/downloads/tobycroft/gorose-pro/total?color=blue)
<a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=P0R-T6lnM--WHzgvGPnbd58US3IUoDlW&jump_from=webapi">
<img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="gorose-orm" title="gorose-orm"></a>

~~~
 ██████╗  ██████╗ ██████╗  ██████╗ ███████╗███████╗    ██████╗ ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗██╔══██╗██╔═══██╗██╔════╝██╔════╝    ██╔══██╗██╔══██╗██╔═══██╗
██║  ███╗██║   ██║██████╔╝██║   ██║███████╗█████╗█████╗██████╔╝██████╔╝██║   ██║
██║   ██║██║   ██║██╔══██╗██║   ██║╚════██║██╔══╝╚════╝██╔═══╝ ██╔══██╗██║   ██║
╚██████╔╝╚██████╔╝██║  ██║╚██████╔╝███████║███████╗    ██║     ██║  ██║╚██████╔╝
 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝    ╚═╝     ╚═╝  ╚═╝ ╚═════╝ 
~~~

## EnglishDOC

[English Document](./README_en.md)

## 原版和Pro版本区别（原版没有的功能）+（猜你关心）

- 反馈群：94537310
- 100%兼容原版
- 本项目已经包含所有Gorose的更新以及Bug修复以及Issues中提到但未修复的问题
- 更加适合ThinkPHP/Laravel开发人员
- go get -u 直接升级，每次升级均做到向上向下兼容无需担心更新后不兼容导致的事故
- 更深度支持MySQL和MariaDB
- 详细文档支持
- 更快的PR/BUG响应+修复速度
- 所有的更新/Bug修复完全来自于当前正在编写的商业项目，不可能出现更新后不管的情况
- (*Pro)支持事务嵌套
- (*Pro)支持复杂Where/OrWhere条件下的and/or条件查询（复杂环境下极好用！）
- (*Pro)CountGroup使用GroupBy的时候返回正确的行数
- (*Pro)SubQuery，安全链式参数化查询操作无需编写语句,生成From *subquery语句*
- (*Pro)SubWhere，安全链式参数化子查询，生成Where *field* *in/=/like...* *subquery*
- (*Pro)修复原版Paginator会出现函数不正确的BUG，高效不出错
- (*Pro)PagiantorWG高性能多线程分页器[性能指示](./doc/performance/PaginatorWG.md)
- (*Pro)修复Executor可能导致故障或删除据的问题
- (*Pro)Oracle数据库支持Replace()方法

## 为什么要使用本项目？

- 项目支持周期`2021-10`~`2028-8`
- 费用：本项目完全免费，劳烦Star
- 本项目已用在金融支付商城教育等项目中，以及GOV项目
- 目前我的项目没有因为GorosePro炸过，可以放心使用
- 立项原因：`原版`商项开发时缺失很多功能，且已`无人维护`
- 原版事务死局：事务在跨模块调用时异常繁琐且没有多级/分级回退功能，这将导致如果你的程序需要设计订单支付功能，在这里有很大的坑等着你
- 原版在实现复杂需求时的代码冗余度非常高，原因是原版更符合面向过程式的开发环境，Pro版本同时支持面向过程和面向对象

## 故障修复

- 修复了高并发下，where等参数的的脏数据问题(如果你在用原版，避免生产环境使用单db)
- 修复了Paginate不能用的问题，并且新增Paginator，让返回更加清晰
- 修复原版Count和GroupBy同时使用时会出现的Total(总条数)错误的问题
- 修复原版Oracle不可用问题，替换驱动使M1以后的ARM芯片可直连

## 商业项目



## 实例文档（Wiki）

- 增删改
    - [增加Insert](../../wiki/Insert新增数据)
    - [删除Delete](../../wiki/Delete删除数据)
    - [更新Update](../../wiki/Update方法)
- 单条查询（对象）(Map[string]interface{})
    - [Find/First返回对象](../../wiki/Find-First查询返回Obj对象方法)
- 多条/联合查询（[]Map[string]interface{}
    - [Get/Select返回数组](../../wiki/Get-Select方法)
    - [Join联合查询](../../wiki/Join-Select方法)
    - [Paginator复杂的子查询分页构建](../../wiki/Paginator复杂的子查询分页构建)
- Query方法
    - [Query方法使用原生语句查询](../../wiki/Query方法)
- 嵌套事务
    - [支付环境下复杂的嵌套事务实例](../../wiki/支付环境下复杂的嵌套事务)
- 子查询subQuery
    - [SubSql防注入From子查询](../../wiki/SubQuery安全子查询)
    - [SubWhere防注入where子查询](../../wiki/SubWhere安全子查询)
- 安全性能
    - [Paginator-Performance-by-ChatGPT](../../wiki/Paginator分页查询的性能问题-ChatGPT )

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
- oracle : https://github.com/sijms/go-ora
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

更多配置, 可以配置集群,甚至可以同时配置不同数据库在一个集群中, 数据库会随机选择集群的数据库来完成对应的读写操作,
其中master是写库, slave是读库, 需要自己做好主从复制, 这里只负责读写

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

