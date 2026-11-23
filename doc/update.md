---
title: Update 修改方法
description: 使用 Data + Update 修改数据，Increment / Decrement 增减字段，事务场景须使用注入法。
order: 4
category: 写入
---

# Update 方法

## 1. 非事务场景（创建法）

场景简单、无需事务时，`update` 成功后即可直接返回成功，使用创建法即可。

**修改单字段：**

```go
func Api_update_uname(qq, uname interface{}) bool {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{"qq": qq})
	db.Data(map[string]interface{}{"uname": uname})
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
```

**修改多字段：**

```go
func Api_update_all(qq, uname, password interface{}) bool {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{"qq": qq})
	db.Data(map[string]interface{}{
		"uname":    uname,
		"password": password,
	})
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
```

## 2. 事务场景（注入法）

资金类修改必须使用注入法：

```go
type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_update(group_id, user_id, balance interface{}) bool {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	})
	db.Data(map[string]interface{}{"balance": balance})
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
```

> **规则**：一旦涉及 Transaction，必须使用注入法；涉及嵌套事务，也必须使用注入法，
> 否则事务与事务隔离均无法生效。

## 3. Increment / Decrement

递增（增量可自定义或传入）：

```go
func (self *Interface) Api_incr(group_id, user_id, cid, amount interface{}) bool {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
	})
	_, err := db.Increment("amount", amount)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
```

递减（也可用负值 + `Increment` 实现递减，注意先 ABS 再取负）：

```go
func (self *Interface) Api_decr(group_id, user_id, dj_id interface{}) bool {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	})
	_, err := db.Decrement("num", 1)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
```

> ORM 已将这些值统一处理为 `interface{}`，只要不传入特殊类型（如 string），一般无需担心类型问题。

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)