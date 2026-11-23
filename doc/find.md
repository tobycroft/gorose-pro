---
title: Find / First 单条查询
description: 使用 First 查询单条数据、Value 取单值，以及创建法与注入法的组合使用。
order: 3
category: 查询
---

# Find / First 方法（单条查询）

## 1. 创建法与注入法组合

大多数场景使用创建法，资金类查询使用注入法。可先写好注入法版本，再写一个创建法版本，
创建后注入 `Interface` 即可：

```go
type Interface struct {
	Db gorose.IOrm
}

func Api_find(group_id, user_id, dj_id interface{}) gorose.Data {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_find(group_id, user_id, dj_id)
}

func (self *Interface) Api_find(group_id, user_id, dj_id interface{}) gorose.Data {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	})
	data, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return data
}
```

## 2. Value 取单值

`find` 之外还有 `value`、`sum` 等方法，大同小异：

```go
func Api_value_balance(group_id, user_id interface{}) interface{} {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	})
	ret, err := db.Value("balance")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return ret
}
```

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)