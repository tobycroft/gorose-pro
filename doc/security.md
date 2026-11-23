---
title: 安全相关（SQL 注入防护）
description: 使用参数化查询防止 SQL 注入，Where 与手写 SQL 的安全用法与危险用法。
order: 8
category: 安全
---

# 安全相关（SQL 注入防护）

## 核心原则

Gorose Pro 已默认使用 prepare（参数化查询），正常情况下不会发生注入。风险主要出现在
**手写 SQL** 或 **不当的 Where 传参** 时。

## 1. 手写 SQL：必须使用参数占位符

**危险（直接拼接参数，会降级为 emulate 模式并导致注入）：**

```go
db.Query("select * from table where user=" + 参数)
```

**安全（使用 `?` 占位符 + 参数）：**

```go
db.Query("select * from table where user=?", 参数)
```

> `Query` 签名为 `Query(sqlstring string, args ...interface{}) ([]Data, error)`，
> 参数应作为可变参数传入，而非拼接到 SQL 字符串中。

## 2. Where 传参：禁止拼接

**危险（字段与值拼接）：**

```go
db.Where("user=" + 参数)
```

**安全（字段与值分离）：**

```go
db.Where("user", "=", 参数)
// 或
db.Where("user", 参数)
```

## 小结

安全无小事，请始终使用参数化查询。数据库层面的安全要点即上述内容，
其他安全细节可参考 ThinkPHP 安全说明。

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)