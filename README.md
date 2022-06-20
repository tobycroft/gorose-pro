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
- 项目文档示例支援更丰富
- 采用"直觉编程"优化，即使没用过也能更快上手
- 100%兼容原版
- 更快的PR/BUG响应+修复速度
- 所有的更新/Bug修复完全来自于当前正在编写的商业项目，不可能出现更新后不管的情况

## 本项目目的

- 为了解决原版框架不再更新维护后可能带来的风险问题
- 为了解决原框架在商业项目实战中出现的各类弊端

## 故障修复

- 修复了高并发下，where等参数的的脏数据问题(如果你在用原版，避免生产环境使用单db)
- 修复了Paginate不能用的问题，并且新增Paginator，让返回更加清晰

## 文档

如下文档均采用解耦模式编写，MAC架构，当然你依然可以使用原文档提供的写法

[文档开发实例1](./doc/intro.md)

## 简介

GorosePro是一个升级改版项目，在支持原框架所有功能的基础上修复了BUG，更加适合复杂的商业项目

支持解耦式开发和直觉式编程，大大降低你的试错成本，让小型项目开发速度比PHP更快，让大型项目更加容易维护，

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

## api预览(详情请参阅文档，或如下演示)

[文档开发实例3](./doc/intro.md)

```go
db.Table("table_name").Fields().Where().GroupBy().Having().OrderBy().Limit().Select()
db.Table(&ModelStruct).Data().Replace()
db.Table(&ModelStruct).Data().Insert()
db.Table(....).Data().Where().Update()
db.Table(....).Where().Delete()
```

## Thinkphp模式用法示例

```go
package main

import (
    "github.com/tobycroft/gorose-pro"
)

func dsn() string {
    dbname := "GobotQ2"
    dbuser := "GobotQ"
    dbpass := "123456"
    dbhost := "10.0.0.170"
    conntype := "tcp"
    dbport := "3306"
    charset := "utf8mb4"
    return dbuser + ":" + dbpass + "@" + conntype + "(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
}

func DbConfig() *gorose.Config {
    var conf gorose.Config
    conf.Driver = "mysql"
    conf.SetMaxIdleConns = 90
    conf.SetMaxOpenConns = 300
    conf.Prefix = ""
    conf.Dsn = dsn()
    return &conf
}

func init() {
    var err error
    Database, err = gorose.Open(DbConfig())
    if err != nil {
        log.Panic(err)
    }
}

func DB() gorose.IOrm {
    return database.Database.NewOrm()
}

//这里是Model层，Model采用单例模式

//增
func Api_insert(qq, token, ip interface{}) bool {
    db := tuuz.Db().Table(table)
    data := map[string]interface{}{
        "qq":    qq,
        "token": token,
        "ip":    ip,
    }
    db.Data(data)
    _, err := db.Insert()
    //_, err := db.Replace()也可以使用replace方法，看你个人
    if err != nil {
        Log.Dbrr(err, tuuz.FUNCTION_ALL())
        return false
    } else {
        return true
    }
}

//删
func Api_delete_byToken(qq, token interface{}) bool {
    db := tuuz.Db().Table(table)
    where := map[string]interface{}{
        "qq":    qq,
        "token": token,
    }
    db.Where(where)
    _, err := db.Delete()
    if err != nil {
        Log.Dbrr(err, tuuz.FUNCTION_ALL())
        return false
    } else {
        return true
    }
}

//修改
func Api_update_password(qq, password interface{}) bool {
    db := tuuz.Db().Table(table)
    where := map[string]interface{}{
        "qq": qq,
    }
    db.Where(where)
    data := map[string]interface{}{
        "password": password,
    }
    db.Data(data)
    _, err := db.Update()
    if err != nil {
        Log.Dbrr(err, tuuz.FUNCTION_ALL())
        return false
    } else {
        return true
    }
}

//查询单条
func Api_find(qq interface{}) gorose.Data {
    db := tuuz.Db().Table(table)
    where := map[string]interface{}{
        "qq": qq,
    }
    db.Where(where)
    ret, err := db.First()
    if err != nil {
        Log.Dbrr(err, tuuz.FUNCTION_ALL())
        return nil
    } else {
        return ret
    }
}

//查询多条
func Api_select(qq interface{}) []gorose.Data {
    db := tuuz.Db().Table(table)
    where := map[string]interface{}{
        "qq": qq,
    }
    db.Where(where)
    ret, err := db.Get()
    if err != nil {
        Log.Dbrr(err, tuuz.FUNCTION_ALL())
        return nil
    } else {
        return ret
    }
}

```

## 使用建议

如果你的数据返回处理比较复杂，并且是“Long Term”项目，这里建议用原版Gorose方法处理，因为我大多数是外包项目， Thinkphp类似的操作方法可以大大降低编码复杂性，性能上并不会差太多，请放心

单例模式极好理解，你可以使用我的方式来解耦，或者也可以使用你自己喜欢的方式

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

这里跳过原生操作，如果你喜欢这么操作，你也不会来用这个框架，这个框架就是简单方便，
自动化没那么多JJYY的规矩，只要你会写Thinkphp或者Laravel，你就可以按照自己的编程习惯来开发，
这一点上我和原作者想法是相同的

## 故障排查

- Gorose存在很多问题，有些问题你可能会遇到，下面列出：
    - 请尽量不要使用框架的主从模式，无论是TP还是Gorose，他能提供的稳定性，一定是不如你直接去买RDS之类的产品的，不要试图在该花钱的时候省钱
    -
  出现锁机制：如果出现锁机制，排查起来请先看慢查询，正常如果时间太长，如果你恰好使用的是我推荐的书写模式，你就能定位超时点，对超时点进行分析即可，老版本在长期使用中确实有出现锁的问题，新版目前没有出现，但是也请大家注意，如果出现了，重启数据库即可解决，如果你对这个功能很不放心，你也可以不使用嵌套查询解决

## Stargazers over time

[![Stargazers over time](https://starchart.cc/tobycroft/gorose-pro.svg)](https://starchart.cc/tobycroft/gorose-pro)
