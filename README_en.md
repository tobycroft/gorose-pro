# GoRose-ORM-Pro A Free Go MySQL ORM

[![GoDoc](https://godoc.org/github.com/tobycroft/gorose-pro?status.svg)](https://godoc.org/github.com/tobycroft/gorose-pro)
[![Go Report Card](https://goreportcard.com/badge/github.com/tobycroft/gorose-pro)](https://goreportcard.com/report/github.com/tobycroft/gorose-pro)
[![GitHub release](https://img.shields.io/github/release/tobycroft/gorose-pro.svg)](https://github.com/tobycroft/gorose-pro/releases/latest)
[![Gitter](https://badges.gitter.im/tobycroft/gorose-pro.svg)](https://gitter.im/gorose-pro/wechat)
![GitHub](https://img.shields.io/github/license/tobycroft/gorose-pro?color=blue)
![GitHub All Releases](https://img.shields.io/github/downloads/tobycroft/gorose-pro/total?color=blue)
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

## What the difference between this than the OriginalVer.

- All functionality was updated to the latest version from the original version
- Intuition Coding
- Nested Transactions functionality
- Simple update without any hesitation ,fully compatible with earlier versions
- Fully support MySQL and MariaDB
- Document support as detail as possible
- 100% compatible with original function
- Dealing PR & Bug fix much more sooner
- All function has been tested and verified by commercial projects
- Support complexity and/or sql in Where/OrWhere function
- Paginator is now able to return the correct TOTAL number
- Add CountGroup to return correct row count from sql sentence
- SubQuery: Highly security parameterized query under prepared condition
- SubWhere: Full Prepared condition parameterized where sql searching
- Fixed original framework's Executor might have unintentionally error or deleted data
- Oracle support "Replace()" function like MyBatis when a good performance

## Purpose of this project

- Fee: Totally Free but only need your star
- Avoid the risk of the deprecation of the original ver
- To solve the shortage during coding in the real life
- Massive demos from the document, what ever the skill you are, you still able to find a solution here

## Bug Fix

- Dirty Read under concurrency circumstances(this will be only and easily triggered by using *db mode)
- Paginate fixed, this function finally come back to life, new "Paginator" function make it much more easier to use
- Fix the row_count(Total in Paginator mode) when using GroupBy function
- Fix Oracle unable to connect & use problem, Fixed arm MacM1 chip's compatibility with Oracle

## Docs and Demos（Wiki）

- CUD functions
    - [Insert](../../wiki/Insert新增数据)
    - [Delete](../../wiki/Delete删除数据)
    - [Update](../../wiki/Update方法)
    - [Replace](../../wiki/Replace方法)
- Read the first line data and return in map[string]any
    - [Find/First](../../wiki/Find-First查询返回Obj对象方法)
- Read multiple data in array by []map[string]any
    - [Get/Select](../../wiki/Get-Select方法)
    - [Join](../../wiki/Join-Select方法)
    - [Paginator](../../wiki/Paginator复杂的子查询分页构建)
    - [Paginator in HighPerformant](../../wiki/PaginatorWG多线程分页)
- Raw SQL sentence mode
    - [Query](../../wiki/Query方法)
- Nested Transaction(only support in GorosePro)
    - [Demos](../../wiki/支付环境下复杂的嵌套事务)
- subQuery
    - [SubSql](../../wiki/SubQuery安全子查询)
    - [SubWhere](../../wiki/SubWhere安全子查询)
- Security and Performance
    - [Paginator-Performance-by-ChatGPT](../../wiki/Paginator分页查询的性能问题-ChatGPT )

## Introduction

Gorosepro is an upgrade and revision project of GOORM. It fixes bugs on the basis of supporting all functions of the
original
framework and is more suitable for complex commercial projects

Support decoupling development and intuitive programming, greatly reduce your trial and error cost, make small projects
develop faster, and make large projects easier to maintain

## Installation

- Add in go.mod

```bash
require github.com/tobycroft/gorose-pro v1.2.5
```

- go get

```bash
go get -u github.com/tobycroft/gorose-pro
```

## Driver support

- mysql : https://github.com/go-sql-driver/mysql
- sqlite3 : https://github.com/mattn/go-sqlite3
- postgres : https://github.com/lib/pq
- oracle : https://github.com/sijms/go-ora
- mssql : https://github.com/denisenkom/go-mssqldb
- clickhouse : https://github.com/kshvakov/clickhouse

## Initializing

- [MySQL](../../wiki/初始化方法MySQL)
- [Oracle](../../wiki/初始化方法Oracle)


For more configurations, you can configure the cluster, or even configure different databases at the same time. In a
cluster, the database will randomly select the database of the cluster to complete the corresponding read-write
operations. The master is the write database, and the slave is the read database. You need to do master-slave
replication. Here, you are only responsible for reading and writing

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

Initialize then use

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
