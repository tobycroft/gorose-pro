---
title: 命名规范与设计思想
description: Gorose Pro 推荐的过程式命名规范（Api_ 前缀、蛇形分割动作 + 驼峰命名）及其设计动机与示例。
order: 0
category: 概览
---

# 命名规范与设计思想

## 1. 为什么方法前缀叫 `Api`

受 Go 访问控制限制，只有首字母大写的标识符才能跨包调用。数据库层采用**单例模式**设计，
所有对数据的调用都以 `Api` 形式对外暴露，因此方法统一以 `Api` 开头。

> 前缀名称可自定义，只要保证跨包可调用即可。

## 2. 命名风格：蛇形分割动作 + 驼峰命名

蛇形（下划线）的作用是**分割动作**，命名主体仍遵循标准驼峰风格。下划线只是表象，驼峰才是本质。

推荐使用「过程」代替「目的」来命名方法，让接手者能快速理解方法行为。

### 命名结构

```
Api_{动作}_{限定条件}
```

| 片段 | 含义 | 示例 |
|------|------|------|
| `Api_` | API 前缀（单例模式） | — |
| `{动作}` | 操作类型（增删改查 / 取值 / 聚合） | `find`、`select`、`value`、`sum`、`update` |
| `{限定条件}` | 查询或修改的限定字段 | `byQqandPassword`、`canShow`、`num` |

### 示例 1：`Api_find_byQqandPassword`

通过方法名即可回答四个问题：

1. 是数据库方法还是逻辑方法？→ 数据库方法（`Api_` 前缀）。
2. 属于增删改查哪一类？→ 查询（`find` = 单条查询）。
3. 返回单条还是多条？用 `map[string]interface{}` 还是 `[]map[string]interface{}` 处理？→ 单条，即 `gorose.Data`。
4. 需要哪两种数据类型完成查询？→ `qq` 与 `password`。

```go
func Api_find_byQqandPassword(qq, password interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where(map[string]interface{}{
		"qq":       qq,
		"password": password,
	})
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return ret
}
```

### 示例 2：`Api_select_canShow`（读取可展示字段）

`select` 表示多条查询，`canShow` 表示读取「可展示」字段。字段名采用**前小后大**（`can` + `Show`）。

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

### 示例 3：`Api_value_num`（复杂限定条件）

当限定条件较多时，可省略 `by` 前缀，直接用「动作 + 目标字段」命名。
**注意**：不要把全部查询限定项都写进方法名，应掌握好度，保持可读性。

```go
func (self *Interface) Api_value_num(group_id, user_id, dj_id interface{}) int64 {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	})
	data, err := db.Value("num")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	}
	if data == nil {
		return 0
	}
	return data.(int64)
}
```

### 示例 4：`Api_sum_byCid`（聚合）

```go
func (self *Interface) Api_sum_byCid(cid interface{}) float64 {
	db := self.Db.Table(table)
	db.Where(map[string]interface{}{"cid": cid})
	ret, err := db.Sum("amount")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	}
	if ret != nil {
		return ret.(float64)
	}
	return 0
}
```

> `tuuz.Db()` 是作者框架中创建数据库对象的方法，读者可按自身项目情况替换。

## 目录

0. [基础准备](./base.md)
1. [Select 方法](./select.md)
2. [Find/First 方法](./find.md)
3. [Update 方法](./update.md)
4. [Delete 方法](./delete.md)
5. [安全相关](./security.md)