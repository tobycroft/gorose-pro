---
title: Join 联表方法
description: 使用 Join / LeftJoin / RightJoin / CrossJoin 进行联表查询的示例。
order: 6
category: 进阶
---

# Join 方法

## 示例

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

## 联表方法清单

| 方法 | 说明 |
|------|------|
| `Join(args ...interface{})` | 内连接（inner join） |
| `LeftJoin(args ...interface{})` | 左连接 |
| `RightJoin(args ...interface{})` | 右连接 |
| `CrossJoin(args ...interface{})` | 交叉连接 |

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)