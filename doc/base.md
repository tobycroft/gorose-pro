---
title: 基础方法：数据库对象创建方式
description: 创建法（每次新建对象）与注入法（外部注入对象）两种数据库对象获取方式，以及事务隔离对选择的影响。
order: 1
category: 基础
---

# 基础方法

以下两种数据库对象获取方式将贯穿所有教程，请务必区分，并根据实际情况选择。

## 1. 创建法（每次查询新建对象）

每次查询都调用 `tuuz.Db()` 创建新的数据库对象：

```go
func Api_find(group_id, user_id, cid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
	})
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return ret
}
```

## 2. 注入法（使用外部注入的对象）

每次查询使用外部注入的数据库对象，注意 `type` 定义：

```go
type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_find(group_id, user_id, cid interface{}) gorose.Data {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
	})
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return ret
}
```

## 3. 如何选择

| 方式 | 适用场景 | 说明 |
|------|---------|------|
| 创建法 | 普通查询 / 修改（推荐） | 每次新建对象，效率最高 |
| 注入法 | 事务（Transaction）/ 嵌套事务 | 必须使用，保证同一数据流 |

### 事务隔离原理

事务是隔离的，不在同一数据流（对象）下的查询无法读取到隔离内的事务修改。

假设数据库字段 `A` 初始值为 `123`：

**无事务场景：**

```
A -> find[输出123] -> update(A,456) -> find[输出456]
```

**find 用创建法、update 用注入法（错误）：**

```
A -> 开始事务 -> (事务外 find[输出123]) -> update(A,456) -> (事务外 find[输出123]) -> 提交事务 -> (事务外 find[输出456])
```

事务外的 `find` 无法读取到未提交的 `456`。

**find / update 均用注入法（正确）：**

```
A -> 开始事务 -> find[输出123] -> update(A,456) -> find[输出456] -> 提交事务 -> find[输出456]
```

> 若开发交易所、资金类程序，必须理解此概念。

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)