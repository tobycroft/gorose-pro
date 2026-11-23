---
title: 编程警告
description: WhereIn 空数组、分页跨页重复排序等常见故障及规避方法。
order: 9
category: 进阶
---

# Gorose Pro 编程警告

## 1. WhereIn 空数组导致 SQL 错误

- **现象**：传入空切片 `[]any{}` 时，语句会变成 `select ... where id in ()`，导致数据库错误。
- **原因**：外部参数可能为空，`WhereIn` 未对空参数做防护。
- **规避**：调用 `WhereIn` 前判断参数是否非空。

## 2. 分页跨页重复（相同排序值）

- **现象**：大量数据分页时，第一页最后一条与第二页前几条重复。
- **原因**：仅按单一字段排序时，排序值相同的记录会随机排列（MySQL 行为）。
- **规避**：排序时加入唯一键（如 `id`）作为次级排序条件。

```go
func Api_paginator_byStudyTypeAndCoinIdAndSchoolIdAndYearAndClassAndDate(coin_id any, extra_like string, school_id, year, class, start_date, end_date any, groupby string, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(table + " a")
	db.Fields("a.*", "b.school_id", "b.year", "b.class", "sum(amount) as sum_amount", "b.name", "c.wx_name", "c.wx_img")
	db.LeftJoin(StudentModel.Table+" b", "a.student_id=b.id")
	db.LeftJoin(UserModel.Table+" c", "a.uid=c.id")
	if coin_id != nil {
		db.Where("coin_id", coin_id)
	}
	if extra_like != "" {
		db.Where("extra", "like", extra_like+".%")
	}
	if school_id != nil {
		db.Where("school_id", school_id)
	}
	if year != nil {
		db.Where("year", year)
	}
	if class != nil {
		db.Where("class", class)
	}
	db.Where("a.date", ">=", start_date)
	db.Where("a.date", "<", end_date)
	db.GroupBy(groupby)
	// 正确：在 sum_amount 基础上加入 id 排行，避免相同积分出现混乱排序
	db.OrderBy("sum_amount desc,id desc")
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Paginator()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	}
	return ret
}
```