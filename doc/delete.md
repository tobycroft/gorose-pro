---
title: Delete 删除方法
description: 使用 Where + Delete 删除数据，注意无 Where 条件时删除会报错。
order: 5
category: 写入
---

# Delete 方法

```go
func Api_delete(qq interface{}) bool {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{"qq": qq})
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
```

> **注意**：没有 `Where` 条件时调用 `Delete()` 会报错（防止误删全表）。

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)