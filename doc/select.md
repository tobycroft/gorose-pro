---
title: Select / Get 多条查询
description: 使用 Where + Get 查询多条数据，范围条件叠加，以及 Join 联表查询。
order: 2
category: 查询
---

# Select 方法（多条查询）

`select` 的核心是 `Where` + `Get`，返回 `[]gorose.Data`（多条结果集）。

## 1. 基础查询

```go
func Api_select(group_id, user_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	})
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return data
}
```

## 2. 叠加范围条件

当 map 形式的 `Where` 无法表达范围查询时，可再次调用 `Where` 的指定模式（`field, op, value`）：

```go
func Api_select_have(group_id, user_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	})
	db.Where("num", ">", 0)
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return data
}
```

## 3. Join 联表查询

```go
func (self *Interface) Api_join_select(group_id, user_id interface{}) []gorose.Data {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	})
	db.Join("coin on coin.id=cid")
	db.Where("amount", ">", 0)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return ret
}
```

> 使用 ORM 时，一对多 / 多对一关系需自行处理 `Join`，但这也让开发者对数据优化有更多认识与提升空间。

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)