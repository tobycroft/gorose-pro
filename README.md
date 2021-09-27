# GoRose-Pro for Commercial(专业版)

## 原版和Pro版本区别（原版没有的功能）+（猜你关心）

- 更加适合ThinkPHP开发人员
- 增加Nested Transaction也就是嵌套事务或者事务嵌套，目前在仿TPORM方面唯一一个支持NT功能的框架
- 缓存支持，原版本功能已经放弃开发，本版在这个部分会使用Redis模块来支持，但是当前功能建议大家使用TuuzGoWeb来解决缓存问题
- 对新特性支持快，replace/nested transaction只要你想它有它就能有
- 去v2后缀，go get -u 直接升级，引入模块无需加v2后缀
- 优先支持MySQL和MariaDB，相信90%用这个框架做商业项目的开发者都是用这两个


[![GoDoc](https://godoc.org/github.com/tobycroft/gorose-pro?status.svg)](https://godoc.org/github.com/tobycroft/gorose-pro)
[![Go Report Card](https://goreportcard.com/badge/github.com/tobycroft/gorose-pro)](https://goreportcard.com/report/github.com/tobycroft/gorose-pro)
[![GitHub release](https://img.shields.io/github/release/tobycroft/gorose.svg)](https://github.com/tobycroft/gorose-pro/releases/latest)
[![Gitter](https://badges.gitter.im/tobycroft/gorose.svg)](https://gitter.im/gorose/wechat)
![GitHub](https://img.shields.io/github/license/tobycroft/gorose?color=blue)
![GitHub All Releases](https://img.shields.io/github/downloads/tobycroft/gorose/total?color=blue)
<a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=P0R-T6lnM--WHzgvGPnbd58US3IUoDlW&jump_from=webapi">
<img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="gorose-orm" title="gorose-orm"></a>

```
 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄               ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄ 
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌             ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀              ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀█░▌
▐░▌          ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌▐░▌          ▐░▌                       ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌
▐░▌ ▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌       ▐░▌
▐░▌▐░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌
▐░▌ ▀▀▀▀▀▀█░▌▐░▌       ▐░▌▐░█▀▀▀▀█░█▀▀ ▐░▌       ▐░▌ ▀▀▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀█░█▀▀ ▐░▌       ▐░▌
▐░▌       ▐░▌▐░▌       ▐░▌▐░▌     ▐░▌  ▐░▌       ▐░▌          ▐░▌▐░▌                       ▐░▌          ▐░▌     ▐░▌  ▐░▌       ▐░▌
▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌      ▐░▌ ▐░█▄▄▄▄▄▄▄█░▌ ▄▄▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄              ▐░▌          ▐░▌      ▐░▌ ▐░█▄▄▄▄▄▄▄█░▌
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌             ▐░▌          ▐░▌       ▐░▌▐░░░░░░░░░░░▌
 ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀               ▀            ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀ 
                                                                                                                                  
```

## 文档

如下的开发实例我已经在自己的项目和多个商业项目中跑过了，代码上没有问题，在书写或者思想上如果和你有冲突你可以用你自己的模式来，这里只是给刚玩的朋友准备的

[文档开发实例1](./doc/intro.md)

[文档开发实例2](./doc/intro.md)



## 简介

gorose for Tuuz版是我从飞哥接手过来的项目，知道人家更新了好几版然后可能有点没兴趣了，但是我是做项目的 很多时候如果原作者不更新，我项目中有很多麻烦事都没办法解决，所以很无奈只能继续扛旗向前走了

因为原版框架已经很优秀了，所以这里只会做一些更新，在架构上不会做出大调整（如果大家满意这个Pro版，请Star）

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

## 对象关系映射, orm的使用

-
    1. 基本链式使用，你可以在测试中这么使用，在做项目时我强烈建议你使用单例模式来调用，
    2. 我对单例模式的支持度时非常深的，这也是Pro版和原版最大的区别
    3. 如果你并不介意repetitive代码，喜欢传统MVC模式开发，请无视如上

```go
var u Users
db := engin.NewOrm()
err := db.Table(&u).Fields("name").AddFields("uid","age").Distinct().Where("uid", ">", 0).OrWhere("age",18).
Group("age").Having("age>1").OrderBy("uid desc").Limit(10).Offset(1).Select()
```

也可以使用`xxx.Limit().Page()`,这个是固定用法,`Page()`必须在`Limit()`后边

-
    2. 如果不想定义struct, 又想绑定指定类型的map结果, 则可以定义map类型, 如

```go
type user gorose.Map
// 或者 以下的type定义, 都是可以正常解析的
type user2 map[string]interface{}
type users3 []user
type users4 []map[string]string
type users5 []gorose.Map
type users6 []gorose.Data
```

- 2.1 开始使用map绑定

```go
db.Table(&user).Select()
db.Table(&users4).Limit(5).Select()
```

> 注意: 如果使用的不是slice数据结构, 则只能获取到一条数据

---
这里使用的 gorose.Data , 实际上就是 `map[string]interface{}` 类型.  
而 `gorose.Map`, 实际上是 `t.MapStringT` 类型, 这里出现了一个 `t` 包, 
是一个golang基本数据类型的相互转换包, 请看详细介绍 http://github.com/gohouse/t

-
    3. laravel的`First()`,`Get()`, 用来返回结果集  
    4. TP使用select和find来取回结果或者结果集，因为Select方法已经被占用，所以请按照TP的Model来理解即可
    5. 你也可以使用直接模式来操作，就是直接填写表名, 返回两个参数, 一个是 `[]gorose.Map`结果集, 第二个是`error`,堪称简单粗暴  
       用法就是把上边的 `Select()` 方法换成 Get,First 即可, 只不过, `Select()` 只返回一个参数
    6. 请不要使用直接表名模式的时候还使用Select方法，取不到数据哦~

-
    7. orm的增删改查

```go
db.Table(&user2).Limit(10.Select()
db.Table(&user2).Where("uid", 1).Data(gorose.Data{"name", "gorose"}).Update()
db.Table(&user2).Data(gorose.Data{"name", "gorose33"}).Insert()
db.Table(&user2).Data([]gorose.Data{{"name", "gorose33"}, "name", "gorose44"}).Insert()
db.Table(&user2).Where("uid", 1).Delete()
```

## 最终sql构造器, builder构造不同数据库的sql

目前支持 mysql, sqlite3, postgres, oracle, mssql, clickhouse等符合 `database/sql` 接口支持的数据库驱动  
这一部分, 用户基本无感知, 分理出来, 主要是为了开发者可以自由添加和修改相关驱动以达到个性化的需求

## binder, 数据绑定对象

这一部分也是用户无感知的, 主要是传入的绑定对象解析和数据绑定, 同样是为了开发者个性化定制而独立出来的

## 模块化

gorose2.0 完全模块化, 每一个模块都封装了interface接口api, 模块间调用, 都是通过接口, 上层依赖下层

- 主模块
    - engin  
      gorose 初始化配置模块, 可以全局保存并复用
    - session  
      真正操作数据库底层模块, 所有的操作, 最终都会走到这里来获取或修改数据
    - orm  
      对象关系映射模块, 所有的orm操作, 都在这里完成
    - builder  
      构建终极执行的sql模块, 可以构建任何数据库的sql, 但要符合`database/sql`包的接口
- 子模块
    - driver  
      数据库驱动模块, 被engin和builder依赖, 根据驱动来搞事情
    - binder  
      结果集绑定模块, 所有的返回结果集都在这里

以上主模块, 都相对独立, 可以个性化定制和替换, 只要实现相应模块的接口即可.




## 故障排查
- Gorose存在很多问题，有些问题你可能会遇到，下面列出：
  - 请尽量不要使用框架的主从模式，无论是TP还是Gorose，他能提供的稳定性，一定是不如你直接去买RDS之类的产品的，不要试图在该花钱的时候省钱
  - 出现锁机制：如果出现锁机制，排查起来请先看慢查询，正常如果时间太长，如果你恰好使用的是我推荐的书写模式，你就能定位超时点，对超时点进行分析即可，老版本在长期使用中确实有出现锁的问题，新版目前没有出现，但是也请大家注意，如果出现了，重启数据库即可解决，如果你对这个功能很不放心，你也可以不使用嵌套查询解决



## Stargazers over time

[![Stargazers over time](https://starchart.cc/tobycroft/gorose-pro.svg)](https://starchart.cc/tobycroft/gorose-pro)
