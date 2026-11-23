---
title: 示例 Model
description: 完整的 Update / First / Get 方法示例，含分步骤注释说明。
order: 7
category: 进阶
---

# 示例 Model

## 修改（Update）

```go
func Api_update_password(qq, password interface{}) bool {
	// 初始化 db 并填写 table 名称
	db := tuuz.Db().Table(table)
	// 创建 where，类型为 map[string]interface{}（与 ThinkPHP 的 Array 相同）
	where := map[string]interface{}{"qq": qq}
	// 导入 where
	db.Where(where)
	// 查询用 db.First() / db.Get() 结束；修改则继续使用 Data
	data := map[string]interface{}{"password": password}
	// 导入 map，使用 Data 方法
	db.Data(data)
	// 执行修改，返回受影响行数与错误
	_, err := db.Update()
	// 判断错误
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
```

## 单条查询（First）

```go
func Api_find(qq interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{"qq": qq})
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return ret
}
```

## 多条查询（Get）

```go
func Api_select_canShow() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{"can_show": 1})
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return data
}
```

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)