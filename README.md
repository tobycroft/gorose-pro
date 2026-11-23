# GoRose-ORM-Pro 完全免费的数据库 ORM

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

> 完全免费 · 100% 兼容原版 · 为复杂商业项目而生

[English Document](./README_en.md) · [版本差异](#版本差异) · [文档目录](#实例文档wiki) · [商业验证](#商业项目验证)

## 版本差异

> `100% 兼容原版` · `支持嵌套事务` · `复杂条件查询` · `高性能分页` · `商业项目验证`

<details>
<summary>展开查看完整差异</summary>

### 兼容与维护

- [x] 100% 兼容原版，支持 `go get -u` 平滑升级，升级前后保持向上、向下兼容
- [x] 已合入原版全部更新、Bug 修复，以及 Issues 中提到但尚未修复的问题
- [x] 本项目仅使用 `v1` 标签；未来如有重构会另开项目，无需担心大版本兼容问题
- [x] 更适合 ThinkPHP / Laravel 开发者，同时支持面向过程与面向对象开发
- [x] 对 MySQL、MariaDB 提供更深入的支持

### Pro 专属能力

- [x] **嵌套事务**：支持跨模块调用时的多级回退
- [x] **复杂 Where / OrWhere**：支持复杂场景下的 `AND` / `OR` 条件组合
- [x] **聚合分页修复**：`CountGroup` 与 `GroupBy` 配合时返回正确行数
- [x] **SubQuery**：安全链式参数化查询，生成 `FROM (subquery)` 语句
- [x] **SubWhere**：安全链式参数化子查询，生成 `WHERE field IN/=/LIKE ... (subquery)`
- [x] **PaginatorWG**：高性能多线程分页器，详见[性能说明](./doc/performance/PaginatorWG.md)
- [x] **Replace()**：Oracle 数据库支持 `Replace()` 方法
- [x] **DataStruct**：避免 `[]Struct` 场景下默认字段导致的 `NULL` 插入或字段丢失
- [x] 修复原版 `Paginator` 返回结果不正确的问题
- [x] 修复原版 `Executor` 可能引发异常或误删数据的问题

</details>

## 商业项目验证

| 数据库 | 已验证版本 / 环境 | 验证状态 |
| --- | --- | :---: |
| MySQL | 全版本；AliYun RDS MySQL 8；BT 8.0.11、8.0.34+ | ✅ |
| MariaDB | 全版本；BT 10.5、10.7；10.10 | ✅ |
| Oracle | 12XE | ✅ |

<details>
<summary>查看项目背景与稳定性说明</summary>

- 项目支持周期：`2021-10` ~ `2028-8`
- 完全免费，只希望你能点个 Star
- 已应用于金融、支付、商城、教育及 GOV 等项目
- 立项原因：原版在商业项目开发中缺失部分关键能力，且已无人维护
- 原版事务在跨模块调用时较为繁琐，缺少多级回退能力，难以支撑订单、支付等复杂场景
- 原版在复杂需求下代码冗余度较高；Pro 版本同时兼顾面向过程与面向对象开发

</details>

## 为什么选择 GoRose-ORM-Pro？

<details>
<summary>展开查看</summary>

GorosePro 是 Gorose ORM 的升级改版项目，在完整支持原框架功能的基础上修复了多个问题，更适合复杂的商业项目。

它支持解耦式开发和直觉式编程，可以显著降低试错成本：让小型项目开发更快，让大型项目更容易维护。

</details>

## 快速开始

```bash
go get -u github.com/tobycroft/gorose-pro@latest
```

或在 `go.mod` 中指定版本：

```go
require github.com/tobycroft/gorose-pro v1.18.1
```

## 实例文档（Wiki）

> `*` 表示 GorosePro 独有功能。

<details>
<summary>展开全部文档入口</summary>

### 增删改

- [增加 Insert](../../wiki/Insert新增数据)
- [删除 Delete](../../wiki/Delete删除数据)
- [更新 Update](../../wiki/Update方法)
- *[替换 Replace](../../wiki/Replace方法)

### 查询

- [Find / First 返回对象](../../wiki/Find-First查询返回Obj对象方法)
- [Get / Select 返回数组](../../wiki/Get-Select方法)
- [Join 联合查询](../../wiki/Join-Select方法)
- [Query 原生语句查询](../../wiki/Query方法)
- *[Paginator 复杂子查询分页](../../wiki/Paginator复杂的子查询分页构建)
- *[PaginatorWG 高性能分页](../../wiki/PaginatorWG多线程分页)

### 事务与子查询

- *[支付场景下的复杂嵌套事务](../../wiki/支付环境下复杂的嵌套事务)
- *[SubQuery 防注入 FROM 子查询](../../wiki/SubQuery安全子查询)
- *[SubWhere 防注入 WHERE 子查询](../../wiki/SubWhere安全子查询)

### Struct 与性能

- *[Scan 将结果输出到 Struct](../../wiki/Scan将结果独立输出到struct)
- [Chunk 分块读取](../../wiki/Chunk分块数据读取)
- *[ChunkWG 高性能分块读取](../../wiki/Chunk分块数据读取)
- [Paginator 分页性能分析](../../wiki/Paginator分页查询的性能问题-ChatGPT )

</details>

## 故障修复（原版未修）

<details>
<summary>展开查看</summary>

- 修复高并发下 `Where` 等参数脏数据问题；使用原版时应避免在生产环境使用单 DB 模式
- 修复 `Paginate` 不可用问题，并新增 `Paginator` 与多线程 `PaginatorWG`
- 修复 `Count` 与 `GroupBy` 同时使用时总条数错误的问题
- 新增 `Counts`，兼容并修复原版复杂场景下计数不准确的问题
- 替换 Oracle 驱动，修复原版 Oracle 不可用问题，并支持 M1 之后的 ARM 芯片直连

</details>

## 支持的驱动

<details>
<summary>展开查看驱动列表</summary>

| 驱动 | 项目地址 |
| --- | --- |
| MySQL | https://github.com/go-sql-driver/mysql |
| SQLite3 | https://github.com/mattn/go-sqlite3 |
| PostgreSQL | https://github.com/lib/pq |
| Oracle | https://github.com/sijms/go-ora |
| MSSQL | https://github.com/denisenkom/go-mssqldb |
| ClickHouse | https://github.com/kshvakov/clickhouse |

</details>

## 配置与初始化

<details>
<summary>展开查看集群配置示例</summary>

- [MySQL 初始化方式](../../wiki/初始化方法MySQL)
- [Oracle 初始化方式](../../wiki/初始化方法Oracle)

可以配置集群，甚至在同一个集群中混用不同数据库。集群会随机选择节点完成读写操作，其中 `Master` 为写库，`Slave` 为读库；主从复制需自行处理，本项目只负责读写路由。

```go
var config1 = gorose.Config{Dsn: "dsn1"}
var config2 = gorose.Config{Dsn: "dsn2"}
var config3 = gorose.Config{Dsn: "dsn3"}
var config4 = gorose.Config{Dsn: "dsn4"}

var configCluster = &gorose.ConfigCluster{
    Master: []gorose.Config{config3, config4},
    Slave:  []gorose.Config{config1, config2},
    Driver: "sqlite3",
}
```

```go
var engin *gorose.Engin
engin, err := Open(config)
// engin, err := Open(configCluster)

if err != nil {
    panic(err.Error())
}
```

</details>

## TODO

<details>
<summary>展开查看</summary>

- 加入缓存层
  - 如果不使用 Struct 接收返回值，Redis Hash 返回的 Column 类型可能统一变成 string
  - 开启缓存后避免类型断言是较简单的方案；若将类型也存入 Redis，虽可解决问题，但会增加 Redis 连接负载
- 可选泛型返回（已完成）
- ElasticSearch
  - 数据自动上载
  - 从 ES 中读取数据

</details>

## Stargazers over time

<details>
<summary>展开查看 Star 趋势</summary>

[![Stargazers over time](https://starchart.cc/tobycroft/gorose-pro.svg?variant=adaptive)](https://starchart.cc/tobycroft/gorose-pro)

</details>
