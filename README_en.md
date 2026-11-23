# GoRose-ORM-Pro A Free Go ORM

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

> Totally free · 100% compatible with the original · Built for complex commercial projects

[中文文档](./README.md) · [Version Differences](#version-differences) · [Documentation](#documentation-wiki) · [Commercial Validation](#commercial-validation)

## Version Differences

> `100% compatible` · `nested transactions` · `complex queries` · `high-performance pagination` · `commercially validated`

<details>
<summary>Expand for full differences</summary>

### Compatibility & maintenance

- [x] 100% compatible with the original version; upgrade smoothly with `go get -u` while keeping upward and downward compatibility
- [x] Includes all updates and bug fixes from the original project, plus issues reported but not yet fixed there
- [x] Uses `v1` tags only; future refactors will be released as separate projects, so no major-version compatibility concerns
- [x] Friendly to ThinkPHP / Laravel developers, supporting both procedural and object-oriented development
- [x] Deeper support for MySQL and MariaDB

### Pro-only capabilities

- [x] **Nested transactions**: multi-level rollback across module boundaries
- [x] **Complex Where / OrWhere**: reliable `AND` / `OR` grouping for complex conditions
- [x] **Correct aggregate pagination**: `CountGroup` returns the correct row count with `GroupBy`
- [x] **SubQuery**: safe, chainable, parameterized `FROM (subquery)` queries
- [x] **SubWhere**: safe, chainable, parameterized `WHERE field IN/=/LIKE ... (subquery)` queries
- [x] **PaginatorWG**: high-performance multithreaded paginator, see [performance notes](./doc/performance/PaginatorWG.md)
- [x] **Replace()**: Oracle support for the `Replace()` method
- [x] **DataStruct**: prevents `NULL` inserts or missing columns caused by default values in `[]Struct`
- [x] Fixes incorrect results returned by the original `Paginator`
- [x] Fixes issues where the original `Executor` could fail or delete data unexpectedly

</details>

## Commercial Validation

| Database | Validated versions / environments | Status |
| --- | --- | :---: |
| MySQL | All versions; AliYun RDS MySQL 8; BT 8.0.11, 8.0.34+ | ✅ |
| MariaDB | All versions; BT 10.5, 10.7; 10.10 | ✅ |
| Oracle | 12XE | ✅ |

<details>
<summary>View project background and stability notes</summary>

- Project support window: `2021-10` ~ `2028-8`
- Totally free — a Star is all we ask for
- Used in finance, payment, e-commerce, education, and government projects
- Created because the original version lacked features needed by commercial projects and is no longer maintained
- The original transaction design is cumbersome across modules and lacks multi-level rollback, making it risky for order and payment workflows
- The original version leads to verbose code for complex requirements; Pro supports both procedural and object-oriented styles

</details>

## Why GoRose-ORM-Pro?

<details>
<summary>Expand to learn more</summary>

GorosePro is an upgraded revision of Gorose ORM. It fixes multiple issues while retaining full support for the original framework, making it more suitable for complex commercial projects.

It supports decoupled development and intuitive programming, significantly reducing trial-and-error costs: smaller projects can move faster, while larger projects stay easier to maintain.

</details>

## Quick Start

```bash
go get -u github.com/tobycroft/gorose-pro@latest
```

Or pin a version in `go.mod`:

```go
require github.com/tobycroft/gorose-pro v1.18.1
```

## Documentation (Wiki)

> `*` indicates a GorosePro-only feature.

<details>
<summary>Expand all documentation links</summary>

### Create, update, delete

- [Insert](../../wiki/Insert新增数据)
- [Delete](../../wiki/Delete删除数据)
- [Update](../../wiki/Update方法)
- *[Replace](../../wiki/Replace方法)

### Query

- [Find / First: return one object](../../wiki/Find-First查询返回Obj对象方法)
- [Get / Select: return a slice](../../wiki/Get-Select方法)
- [Join](../../wiki/Join-Select方法)
- [Query: raw SQL](../../wiki/Query方法)
- *[Paginator: complex subquery pagination](../../wiki/Paginator复杂的子查询分页构建)
- *[PaginatorWG: high-performance pagination](../../wiki/PaginatorWG多线程分页)

### Transactions & subqueries

- *[Complex nested transactions in payment scenarios](../../wiki/支付环境下复杂的嵌套事务)
- *[SubQuery: injection-safe FROM subquery](../../wiki/SubQuery安全子查询)
- *[SubWhere: injection-safe WHERE subquery](../../wiki/SubWhere安全子查询)

### Structs & performance

- *[Scan results into a struct](../../wiki/Scan将结果独立输出到struct)
- [Chunk: read data in batches](../../wiki/Chunk分块数据读取)
- *[ChunkWG: high-performance batched reads](../../wiki/Chunk分块数据读取)
- [Paginator performance analysis](../../wiki/Paginator分页查询的性能问题-ChatGPT )

</details>

## Bug Fixes (not fixed in the original)

<details>
<summary>Expand to view</summary>

- Fixed dirty parameter data for `Where` and related operations under high concurrency; avoid single-DB mode in production with the original version
- Fixed the unusable `Paginate` function and added `Paginator` plus the multithreaded `PaginatorWG`
- Fixed incorrect totals when `Count` and `GroupBy` are used together
- Added `Counts` for compatibility and fixed inaccurate counts in complex scenarios in the original version
- Replaced the Oracle driver to fix the original Oracle connectivity issue and support direct connections on M1 and later ARM chips

</details>

## Driver Support

<details>
<summary>Expand to view supported drivers</summary>

| Driver | Repository |
| --- | --- |
| MySQL | https://github.com/go-sql-driver/mysql |
| SQLite3 | https://github.com/mattn/go-sqlite3 |
| PostgreSQL | https://github.com/lib/pq |
| Oracle | https://github.com/sijms/go-ora |
| MSSQL | https://github.com/denisenkom/go-mssqldb |
| ClickHouse | https://github.com/kshvakov/clickhouse |

</details>

## Configuration & Initialization

<details>
<summary>Expand for cluster configuration examples</summary>

- [MySQL initialization](../../wiki/初始化方法MySQL)
- [Oracle initialization](../../wiki/初始化方法Oracle)

Clusters can be configured with different databases in the same cluster. A cluster randomly selects a node for read or write operations: `Master` nodes handle writes and `Slave` nodes handle reads. Master-slave replication must be configured separately; GorosePro only handles read/write routing.

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
<summary>Expand to view</summary>

- Add a caching layer
  - Without a Struct receiver, Redis Hash column types may all become strings
  - Avoiding type assertions after caching is enabled is the simpler approach; storing types in Redis also works but increases Redis connection load
- Optional generic returns (done)
- ElasticSearch
  - Automatic data upload
  - Read data from ES

</details>

## Stargazers over time

<details>
<summary>Expand to view the star chart</summary>

[![Stargazers over time](https://starchart.cc/tobycroft/gorose-pro.svg)](https://starchart.cc/tobycroft/gorose-pro)

</details>
